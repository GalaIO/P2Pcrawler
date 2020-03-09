package peerwire

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/GalaIO/P2Pcrawler/misc"
	"io/ioutil"
	"net"
	"strings"
	"time"
)

// refer http://www.bittorrent.org/beps/bep_0009.html
const SizeOf16KB = 16 * 1024

type ExFecthMetaType int

const (
	ExRequest ExFecthMetaType = iota
	ExData
	ExReject
)

type ExtendedFetchMetaMsg interface {
	ExMessageId() int
	MsgType() ExFecthMetaType
	PieceNum() int
	Data() []byte
	TotalSize() int
}

type BaseExFetchMetaMsg struct {
	msgId     int
	msgType   ExFecthMetaType
	pieceNum  int
	data      []byte
	totalSize int
}

func (b *BaseExFetchMetaMsg) TotalSize() int {
	return b.totalSize
}

func NewBaseExFetchMetaMsg(msgId int, msgType ExFecthMetaType, pieceNum, totalSize int, data []byte) *BaseExFetchMetaMsg {
	return &BaseExFetchMetaMsg{msgId: msgId, msgType: msgType, pieceNum: pieceNum, totalSize: totalSize, data: data}
}

func (b *BaseExFetchMetaMsg) ExMessageId() int {
	return b.msgId
}

func (b *BaseExFetchMetaMsg) MsgType() ExFecthMetaType {
	return b.msgType
}

func (b *BaseExFetchMetaMsg) PieceNum() int {
	return b.pieceNum
}

func (b *BaseExFetchMetaMsg) Data() []byte {
	return b.data
}

func (b *BaseExFetchMetaMsg) Bytes() []byte {
	dict := misc.Dict{"msg_type": int(b.msgType), "piece": b.pieceNum}
	if len(b.data) > 0 {
		dict["total_size"] = b.totalSize
	}
	dst, err := misc.EncodeDict(dict)
	if err != nil {
		peerWireLogger.Panic("encode extended fetchmetadata err", misc.Dict{"msgtype": b.msgType, "err": err})
	}

	preLen := 2 + len(dst) + len(b.data)
	preLenBytes := make([]byte, PrefixLen)
	binary.BigEndian.PutUint32(preLenBytes, uint32(preLen))

	var buf bytes.Buffer
	buf.Grow(preLen + PrefixLen)
	buf.Write(preLenBytes)
	buf.WriteByte(byte(ExtendedPeerMsg)) // peer msg type, extended msg
	buf.WriteByte(byte(b.msgId))
	buf.Write([]byte(dst))
	if len(b.data) > 0 {
		buf.Write(b.data)
	}
	return buf.Bytes()
}

func parseExtendedFetchMetaMsg(data []byte) ExtendedFetchMetaMsg {
	fmt.Println("parseExtendedFetchMetaMsg", hex.EncodeToString(data))
	msgType := PeerMsgType(data[0])
	if ExtendedPeerMsg != msgType {
		peerWireLogger.Panic("extended fetchmetadata not extended msg", misc.Dict{"totalLen": len(data)})
	}
	msgId := int(data[1])
	dict, next, err := misc.DecodeDictNoLimit(misc.Bytes2Str(data[2:]))
	next += 2
	if err != nil {
		peerWireLogger.Panic("decode extended fetchmetadata err", misc.Dict{"err": err})
	}
	if next >= len(data) {
		return NewBaseExFetchMetaMsg(msgId, ExFecthMetaType(dict.GetInteger("msg_type")), dict.GetInteger("piece"), 0, nil)
	}

	totalSize := dict.GetInteger("total_size")
	return NewBaseExFetchMetaMsg(msgId, ExFecthMetaType(dict.GetInteger("msg_type")), dict.GetInteger("piece"), totalSize, data[next:])
}

func withExtendedFetchMetaMsg(exMsgId int, mType ExFecthMetaType, pieceNum, totalSize int, data ...[]byte) []byte {
	var binData []byte
	if len(data) > 0 {
		binData = data[0]
	}
	fetchMetaMsg := NewBaseExFetchMetaMsg(exMsgId, mType, pieceNum, totalSize, binData)
	return fetchMetaMsg.Bytes()
}

var fetchMetaLogger = misc.GetLogger().SetPrefix("FetchMetadata")

// fetch .torrent file from peer
func FetchMetaData(laddr string, infoHash []byte) (ret []byte, retErr error) {
	defer func() {
		if err := recover(); err != nil {
			ret = nil
			fetchMetaLogger.Error("FetchMetaData err", misc.Dict{"laddr": laddr, "err": err})
			retErr = errors.New("FetchMetaData err")
		}
	}()
	conn, err := net.DialTimeout("tcp", laddr, 3*time.Second)
	if err != nil {
		fetchMetaLogger.Panic("connect peer err", misc.Dict{"laddr": laddr, "err": err})
	}

	// handshake, exchange info
	conn.Write(withHandShakeMsg(localPeerId, infoHash))

	// get handshake response
	readBytes, err := ioutil.ReadAll(conn)
	if err != nil {
		fetchMetaLogger.Panic("get handshake response err", misc.Dict{"laddr": laddr, "err": err})
	}
	handShakeResp := parseHandShakeMsg(readBytes)
	if !strings.EqualFold(supportProtocolName, handShakeResp.Protocol()) {
		fetchMetaLogger.Panic("get handshake not support bittoroute protocol", misc.Dict{"laddr": laddr, "handShakeResp": handShakeResp})
	}
	if !handShakeResp.SupportExtended() {
		fetchMetaLogger.Panic("get handshake not support extended protocol", misc.Dict{"laddr": laddr, "handShakeResp": handShakeResp})
	}

	// extended handshake, exchange info
	msgId := 3
	conn.Write(withExtendedhandShake(misc.Dict{"ut_metadata": msgId}, nil))

	// get extended handshake response
	readBytes, err = ioutil.ReadAll(conn)
	if err != nil {
		fetchMetaLogger.Panic("get extended handshake response err", misc.Dict{"laddr": laddr, "err": err})
	}
	exHandShakeResp := parseExtendedHandShake(readBytes)
	metadata := exHandShakeResp.Metadata()
	if msgId != metadata.GetInteger("ut_metadata") {
		fetchMetaLogger.Panic("get extended handshake not support", misc.Dict{"laddr": laddr, "exHandShakeResp": exHandShakeResp})
	}
	dict := exHandShakeResp.Dict()
	if !dict.Exist("metadata_size") {
		fetchMetaLogger.Panic("get extended handshake wrong format", misc.Dict{"laddr": laddr, "exHandShakeResp": exHandShakeResp})
	}

	// get piece
	metaSize := dict.GetInteger("metadata_size")
	pieceCount := metaSize / SizeOf16KB
	if metaSize%SizeOf16KB > 0 {
		pieceCount++
	}
	var fileBytes bytes.Buffer
	fileBytes.Grow(metaSize)
	for i := 0; i < pieceCount; i++ {
		// request piece
		conn.Write(withExtendedFetchMetaMsg(msgId, ExRequest, i, 0))

		// get piece
		bytes, err := ioutil.ReadAll(conn)
		if err != nil {
			fetchMetaLogger.Panic("get extended piece response err", misc.Dict{"laddr": laddr, "err": err})
		}
		fetchMetaResp := parseExtendedFetchMetaMsg(bytes)
		// check if same msgId, not reject msg, and correct piece num
		if ExData != fetchMetaResp.MsgType() || msgId != fetchMetaResp.ExMessageId() || i != fetchMetaResp.PieceNum() {
			fetchMetaLogger.Panic("get extended piece wrong format", misc.Dict{"laddr": laddr, "fetchMetaResp": fetchMetaResp})
		}
		fileBytes.Write(fetchMetaResp.Data())
	}
	return fileBytes.Bytes(), nil
}
