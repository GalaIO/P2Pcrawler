package peerwire

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"github.com/GalaIO/P2Pcrawler/misc"
	"io"
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
	conn, err := net.DialTimeout("tcp", laddr, 3*time.Second)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := recover(); err != nil {
			//debug.PrintStack()
			ret = nil
			fetchMetaLogger.Error("FetchMetaData err", misc.Dict{"laddr": laddr, "err": err})
			retErr = errors.New("FetchMetaData err")
		}
		conn.Close()
	}()
	conn.SetReadDeadline(time.Now().Add(time.Second * 1))
	conn.SetWriteDeadline(time.Time{})
	peerConn := NewPeerConn(laddr, conn, infoHash, time.Second*10)
	handshakePeer(peerId, peerConn)
	initialMessages(peerConn)
	readMetaDataLoop(peerConn)

	// merge pieces
	if !peerConn.CompledMetaData() {
		return nil, errors.New("cannot completed download metadata")
	}

	result := make([]byte, 0, peerConn.metaDataSize)
	for i := 0; i < len(peerConn.requestedMetaData); i++ {
		result = append(result, peerConn.metaData[i]...)
	}
	// checksum
	if !bytes.Equal(peerConn.infoHash, GenerateInfoHash(result)) {
		fetchMetaLogger.Error("chesum metadata not match", misc.Dict{"laddr": laddr, "result": hex.EncodeToString(result)})
		return nil, errors.New("check sum err")
	}
	fetchMetaLogger.Info("chesum metadata match", misc.Dict{"laddr": laddr, "infohash": hex.EncodeToString(peerConn.infoHash)})

	return result, nil
}

// read loop for communication
func readMetaDataLoop(peerConn *PeerConn) {

	laddr := peerConn.addr
	conn := peerConn.conn
	for !peerConn.CompledMetaData() && !peerConn.IsReadLoopExpire() {
		readBytes, err := readBytesByPrefixLenMsg(conn)
		if err == io.EOF || len(readBytes) == 0 {
			// keep liave
			//fetchMetaLogger.Info("keep alive msg", misc.Dict{"laddr": laddr})
			continue
		}
		if err != nil {
			fetchMetaLogger.Panic("read prefix length err", misc.Dict{"laddr": laddr, "err": err})
		}
		prefixLenMsg := parsePrefixLenMsg(readBytes)
		switch prefixLenMsg.PeerMsgType() {
		case ExtendedPeerMsg:
			handleExtendedMsg(peerConn, readBytes)
		default:
			// other msg just pass
			fetchMetaLogger.Info("fetch bep3 msg", misc.Dict{"laddr": laddr, "peerMsgType": int(prefixLenMsg.PeerMsgType()), "payload": hex.EncodeToString(readBytes)})
		}
	}

}

// handle extend msg
func handleExtendedMsg(peerConn *PeerConn, readBytes []byte) {

	laddr := peerConn.addr
	msgId := int(readBytes[1])
	switch msgId {
	//HandshakeExtendedID
	case 0:
		// get extended handshake response
		exHandShakeResp := parseExtendedHandShake(readBytes)
		fetchMetaLogger.Info("get extended handshake", misc.Dict{"laddr": laddr, "exHandShakeResp": exHandShakeResp})

		dict := exHandShakeResp.Dict()
		if !dict.Exist("metadata_size") {
			fetchMetaLogger.Panic("get extended handshake wrong format", misc.Dict{"laddr": laddr, "exHandShakeResp": exHandShakeResp})
		}

		// get piece
		metaSize := dict.GetInteger("metadata_size")
		peerConn.metaDataSize = metaSize
		pieceCount := metaSize / SizeOf16KB
		if metaSize%SizeOf16KB > 0 {
			pieceCount++
		}
		peerConn.metaData = make([][]byte, pieceCount)
		peerConn.requestedMetaData = make([]bool, pieceCount)
		for i := 0; i < pieceCount; i++ {
			peerConn.requestedMetaData[i] = false
			// request piece
			_, err := peerConn.conn.Write(withExtendedFetchMetaMsg(1, ExRequest, i, 0))
			if err != nil {
				fetchMetaLogger.Panic("write request piece err", misc.Dict{"laddr": laddr, "err": err})
			}
		}
		fetchMetaLogger.Info("request pieces", misc.Dict{"laddr": laddr, "metaSize": metaSize, "pieceCount": pieceCount})
	//metadataExtendedId
	case 1:
		fetchMetaResp := parseExtendedFetchMetaMsg(readBytes)
		fetchMetaLogger.Info("fetch extended msg", misc.Dict{"laddr": laddr, "fetchMetaResp": fetchMetaResp})
		// check if same MsgId, not reject msg, and correct piece num
		if ExData != fetchMetaResp.MsgType() {
			fetchMetaLogger.Error("get extended piece wrong format", misc.Dict{"laddr": laddr, "fetchMetaResp": fetchMetaResp})
			return
		}
		pieceNum := fetchMetaResp.PieceNum()
		peerConn.metaData[pieceNum] = fetchMetaResp.Data()
		peerConn.requestedMetaData[pieceNum] = true
	//pexExtendedId
	case 2:
		fetchMetaLogger.Info("fetch pexExtendedId extend msg", misc.Dict{"laddr": laddr, "msgid": msgId, "payload": hex.EncodeToString(readBytes)})
	default:
		fetchMetaLogger.Info("fetch pexExtendedId wrong extend msg", misc.Dict{"laddr": laddr, "msgid": msgId, "payload": hex.EncodeToString(readBytes)})

	}

}

// send bep extend message, ie. extendedHandshake, fast, dht port
func initialMessages(peerConn *PeerConn) {
	// extended handshake, exchange info, ut_metadata = 1 is support extended
	_, err := peerConn.conn.Write(withExtendedhandShake(misc.Dict{"ut_metadata": 1}, nil))
	if err != nil {
		fetchMetaLogger.Panic("write extended handshake err", misc.Dict{"laddr": peerConn.addr, "err": err})
	}
}

// handshake, exchange info
func handshakePeer(peerId []byte, peerConn *PeerConn) {
	laddr := peerConn.addr
	_, err := peerConn.conn.Write(withHandShakeMsg(peerId, peerConn.infoHash))
	if err != nil {
		fetchMetaLogger.Panic("write handshake err", misc.Dict{"laddr": laddr, "err": err})
	}
	// get handshake response
	readBytes := make([]byte, handShakeLen)
	_, err = peerConn.conn.Read(readBytes)
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
}
