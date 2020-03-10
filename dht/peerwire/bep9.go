package peerwire

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/GalaIO/P2Pcrawler/misc"
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
	MsgId       int
	ExMsgType   ExFecthMetaType
	ExPieceNum  int
	data        []byte
	ExTotalSize int
}

func (b *BaseExFetchMetaMsg) TotalSize() int {
	return b.ExTotalSize
}

func NewBaseExFetchMetaMsg(msgId int, msgType ExFecthMetaType, pieceNum, totalSize int, data []byte) *BaseExFetchMetaMsg {
	return &BaseExFetchMetaMsg{MsgId: msgId, ExMsgType: msgType, ExPieceNum: pieceNum, ExTotalSize: totalSize, data: data}
}

func (b *BaseExFetchMetaMsg) ExMessageId() int {
	return b.MsgId
}

func (b *BaseExFetchMetaMsg) MsgType() ExFecthMetaType {
	return b.ExMsgType
}

func (b *BaseExFetchMetaMsg) PieceNum() int {
	return b.ExPieceNum
}

func (b *BaseExFetchMetaMsg) Data() []byte {
	return b.data
}

func (b *BaseExFetchMetaMsg) Bytes() []byte {
	dict := misc.Dict{"msg_type": int(b.ExMsgType), "piece": b.ExPieceNum}
	if len(b.data) > 0 {
		dict["total_size"] = b.ExTotalSize
	}
	dst, err := misc.EncodeDict(dict)
	if err != nil {
		peerWireLogger.Panic("encode extended fetchmetadata err", misc.Dict{"msgtype": b.ExMsgType, "err": err})
	}

	preLen := 2 + len(dst) + len(b.data)
	preLenBytes := make([]byte, PrefixLen)
	binary.BigEndian.PutUint32(preLenBytes, uint32(preLen))

	var buf bytes.Buffer
	buf.Grow(preLen + PrefixLen)
	buf.Write(preLenBytes)
	buf.WriteByte(byte(ExtendedPeerMsg)) // peer msg type, extended msg
	buf.WriteByte(byte(b.MsgId))
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
func FetchMetaData(laddr string, peerId, infoHash []byte) (ret []byte, retErr error) {
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
	_, err = conn.Write(withHandShakeMsg(peerId, infoHash))
	if err != nil {
		fetchMetaLogger.Panic("write handshake err", misc.Dict{"laddr": laddr, "err": err})
	}

	// get handshake response
	readBytes := make([]byte, handShakeLen)
	_, err = conn.Read(readBytes)
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
	msgId := 1
	_, err = conn.Write(withExtendedhandShake(misc.Dict{"ut_metadata": msgId}, nil))
	if err != nil {
		fetchMetaLogger.Panic("write extended handshake err", misc.Dict{"laddr": laddr, "err": err})
	}

	// get extended handshake response
	readBytes, err = readBytesByPrefixLenMsg(conn)
	if err != nil {
		fetchMetaLogger.Panic("get extended handshake response err", misc.Dict{"laddr": laddr, "err": err})
	}
	exHandShakeResp := parseExtendedHandShake(readBytes)
	fetchMetaLogger.Info("get extended handshake", misc.Dict{"laddr": laddr, "exHandShakeResp": exHandShakeResp})

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
	fileBytes := make([][]byte, pieceCount)
	for i := 0; i < pieceCount; i++ {
		// request piece
		_, err := conn.Write(withExtendedFetchMetaMsg(msgId, ExRequest, i, 0))
		if err != nil {
			fetchMetaLogger.Panic("write request piece err", misc.Dict{"laddr": laddr, "err": err})
		}
	}
	totalSize := 0
	count := 0
	// only read pieceCount times, ignore pep msg
	for i := 0; i < pieceCount; i++ {
		bytes, err := readBytesByPrefixLenMsg(conn)
		if err != nil {
			fetchMetaLogger.Panic("get extended piece response err", misc.Dict{"laddr": laddr, "err": err})
		}

		prefixLenMsg := parsePrefixLenMsg(bytes)
		if ExtendedPeerMsg != prefixLenMsg.PeerMsgType() {
			fetchMetaLogger.Info("fetch bep3 msg", misc.Dict{"laddr": laddr, "ExMsgType": int(prefixLenMsg.PeerMsgType())})
			continue
		}
		fetchMetaResp := parseExtendedFetchMetaMsg(bytes)
		fetchMetaLogger.Info("fetch extended msg", misc.Dict{"laddr": laddr, "fetchMetaResp": fetchMetaResp})
		// check if same MsgId, not reject msg, and correct piece num
		if ExData != fetchMetaResp.MsgType() || msgId != fetchMetaResp.ExMessageId() {
			fetchMetaLogger.Error("get extended piece wrong format", misc.Dict{"laddr": laddr, "fetchMetaResp": fetchMetaResp})
			break
		}
		pieceNum := fetchMetaResp.PieceNum()
		fileBytes[pieceNum] = fetchMetaResp.Data()
		totalSize += len(fetchMetaResp.Data())
		count++
	}

	// merge pieces
	result := make([]byte, 0, totalSize)
	for i := 0; i < count; i++ {
		result = append(result, fileBytes[i]...)
	}
	fetchMetaLogger.Info("fetch data", misc.Dict{"result": hex.EncodeToString(result)})
	// checksum
	if !bytes.Equal(infoHash, GenerateInfoHash(result)) {
		fetchMetaLogger.Panic("chesum metadata not match", misc.Dict{"laddr": laddr})
	}
	conn.Close()
	return result, nil
}
