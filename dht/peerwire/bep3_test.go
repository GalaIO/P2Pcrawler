package peerwire

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenHandshakeMsg(t *testing.T) {
	infoHash := generatePeerId("infohash")
	peerId := generatePeerId("peerId")
	data := withHandShakeMsg(peerId, infoHash)
	handShakeMsg := parseHandShakeMsg(data)
	assert.Equal(t, "BitTorrent protocol", handShakeMsg.Protocol())
	assert.Equal(t, true, handShakeMsg.SupportExtended())
	assert.Equal(t, peerId, handShakeMsg.PeerId())
	assert.Equal(t, infoHash, handShakeMsg.InfoHash())
}
