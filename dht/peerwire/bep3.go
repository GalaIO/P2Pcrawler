package peerwire

import (
	"bytes"
	"encoding/binary"
	"github.com/GalaIO/P2Pcrawler/misc"
)

// refer http://www.bittorrent.org/beps/bep_0003.html
const handShakeLen = 68
const PeerIdLen = 20
const InfoHashLen = 20
const PrefixLen = 4

var defaultReservedBytes = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
var supportProtocolName = "BitTorrent protocol"
var reservedBytes = supportExterned(defaultReservedBytes)
var localPeerId = generatePeerId("galaio.peerId")

var peerWireLogger = misc.GetLogger().SetPrefix("peerwire")

type HandShakeMsg interface {
	Protocol() string
	SupportExtended() bool
	PeerId() []byte
	InfoHash() []byte
}

type BaseHandShakeMsg struct {
	protocol string
	rBytes   []byte
	infoHash []byte
	peerId   []byte
}

func NewBaseHandShakeMsg(protocol string, rBytes []byte, infoHash []byte, peerId []byte) *BaseHandShakeMsg {
	return &BaseHandShakeMsg{protocol: protocol, rBytes: rBytes, infoHash: infoHash, peerId: peerId}
}

func (b *BaseHandShakeMsg) Protocol() string {
	return b.protocol
}

func (b *BaseHandShakeMsg) SupportExtended() bool {
	return b.rBytes[5]&0x10 > 0
}

func (b *BaseHandShakeMsg) PeerId() []byte {
	return b.peerId
}

func (b *BaseHandShakeMsg) InfoHash() []byte {
	return b.infoHash
}

func (b *BaseHandShakeMsg) Bytes() []byte {
	var buf bytes.Buffer
	buf.Grow(handShakeLen)
	buf.WriteByte(byte(len(b.protocol)))
	buf.WriteString(b.protocol)
	buf.Write(b.rBytes)
	buf.Write(b.infoHash)
	buf.Write(b.peerId)
	return buf.Bytes()
}

func parseHandShakeMsg(data []byte) HandShakeMsg {
	if handShakeLen != len(data) {
		peerWireLogger.Panic("parseHandShakeMsg wrong data len", misc.Dict{"len": len(data)})
	}
	index := 0
	protocolLen := int(data[index])
	index++
	protocol := data[index : index+protocolLen]
	index += protocolLen
	rBytes := data[index : index+len(reservedBytes)]
	index += len(reservedBytes)
	infoHash := data[index : index+InfoHashLen]
	index += InfoHashLen
	peerId := data[index : index+PeerIdLen]
	return NewBaseHandShakeMsg(string(protocol), rBytes, infoHash, peerId)
}

func withHandShakeMsg(peerId, infoHash []byte) []byte {
	if PeerIdLen != len(peerId) || InfoHashLen != len(infoHash) {
		peerWireLogger.Panic("wrong peerId or infohash", misc.Dict{"peerIdLen": len(peerId), "infohashLen": len(infoHash)})
	}
	handShakeMsg := NewBaseHandShakeMsg(supportProtocolName, reservedBytes, infoHash, peerId)
	return handShakeMsg.Bytes()
}

func supportExterned(bytes []byte) []byte {
	bytes[5] = bytes[5] | 0x10
	return bytes
}

func parsePrefixLen(data []byte) int {
	preLen := binary.BigEndian.Uint32(data[:PrefixLen])
	if PrefixLen+preLen > uint32(len(data)) {
		peerWireLogger.Panic("parseExtendedHandShake wrong data len", misc.Dict{"len": len(data)})
	}
	return int(preLen)
}
