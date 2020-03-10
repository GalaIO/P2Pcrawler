package dht

import (
	"github.com/GalaIO/P2Pcrawler/dht/krpc"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestParseFetchAddr(t *testing.T) {
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8000")
	req := WithAnnouncePeerMsg("aa", "bb", "infohash", "token", 9000, nil)
	assert.Equal(t, "127.0.0.1:9000", parseFetchAddr(krpc.NewReqContext(nil, req, nil, addr)))
}
