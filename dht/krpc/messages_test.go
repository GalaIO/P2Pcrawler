package krpc

import (
	"github.com/GalaIO/P2Pcrawler/dht"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenQueryMsg(t *testing.T) {

	var req Request

	req = withPingMsg("aa", testNodeId, func(req Request, resp Response) {

	})
	assert.Equal(t, dht.Dict{"t": "aa", "y": "q", "q": "ping", "a": dht.Dict{"id": "abcdefghij0123456789"}}, req.RawData())

	req = withFindNodeMsg("aa", "abcdefghij0123456789", "mnopqrstuvwxyz123456", func(req Request, resp Response) {

	})
	assert.Equal(t, dht.Dict{"t": "aa", "y": "q", "q": "find_node", "a": dht.Dict{"id": "abcdefghij0123456789", "target": "mnopqrstuvwxyz123456"}}, req.RawData())

}
