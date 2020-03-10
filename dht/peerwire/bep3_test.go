package peerwire

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenHandshakeMsg(t *testing.T) {
	infoHash, _ := hex.DecodeString("9bd56482d6fd6a436f5051d3b9560cdd942a5962")
	peerId := GeneratePeerId("test1")
	data := withHandShakeMsg(peerId, infoHash)
	assert.Equal(t, "13426974546f7272656e742070726f746f636f6c00000000001000019bd56482d6fd6a436f5051d3b9560cdd942a5962b444ac06613fc8d63795be9ad0beaf55011936ac", hex.EncodeToString(data))
	handShakeMsg := parseHandShakeMsg(data)
	assert.Equal(t, "BitTorrent protocol", handShakeMsg.Protocol())
	assert.Equal(t, true, handShakeMsg.SupportExtended())
	assert.Equal(t, peerId, handShakeMsg.PeerId())
	assert.Equal(t, infoHash, handShakeMsg.InfoHash())
}

func TestParseHandshakeMsg(t *testing.T) {
	infoHash, _ := hex.DecodeString("9bd56482d6fd6a436f5051d3b9560cdd942a5962")
	data, _ := hex.DecodeString("13426974546f7272656e742070726f746f636f6c00000000001000049bd56482d6fd6a436f5051d3b9560cdd942a59622d5350333630355f549fe731afaa6fdb0494e971")
	handShakeMsg := parseHandShakeMsg(data)
	assert.Equal(t, supportProtocolName, handShakeMsg.Protocol())
	assert.Equal(t, true, handShakeMsg.SupportExtended())
	assert.Equal(t, infoHash, handShakeMsg.InfoHash())
}

func TestParsePrefixLenMsg(t *testing.T) {
	data, _ := hex.DecodeString("05000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	preLenMsg := parsePrefixLenMsg(data)
	assert.Equal(t, BitFieldPeerMsg, preLenMsg.PeerMsgType())

	data, _ = hex.DecodeString("02")
	preLenMsg = parsePrefixLenMsg(data)
	assert.Equal(t, BitFieldPeerMsg, preLenMsg.PeerMsgType())

	data, _ = hex.DecodeString("01")
	preLenMsg = parsePrefixLenMsg(data)
	assert.Equal(t, BitFieldPeerMsg, preLenMsg.PeerMsgType())
}
