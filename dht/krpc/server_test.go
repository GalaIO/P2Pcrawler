package krpc

import (
	"github.com/GalaIO/P2Pcrawler/dht"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testNodeId = "abcdefghij0123456789"

func TestHandlePingMsg(t *testing.T) {
	RegisteQueryHandler("ping", func(ctx *QueryCtx) dht.Dict {
		return dht.Dict{"id": "mnopqrstuvwxyz123456"}
	})

	resp := queriesHandle(dht.Dict{"t": "aa", "y": "q", "q": "ping", "a": dht.Dict{"id": "abcdefghij0123456789"}})
	assert.Equal(t, dht.Dict{"t": "aa", "y": "r", "r": dht.Dict{"id": "mnopqrstuvwxyz123456"}}, resp)

	resp = queriesHandle(dht.Dict{"t": "aa", "y": "q", "q": "ping"})
	assert.Equal(t, dht.Dict{"t": "aa", "y": "e", "e": dht.List{dht.ProtocolErr, "201:cannot find a's val"}}, resp)
}

func TestHandleFindNodeMsg(t *testing.T) {
	RegisteQueryHandler("find_node", func(ctx *QueryCtx) dht.Dict {
		return dht.Dict{"id": "0123456789abcdefghij", "nodes": "def456..."}
	})

	resp := queriesHandle(dht.Dict{"t": "aa", "y": "q", "q": "find_node", "a": dht.Dict{"id": "abcdefghij0123456789", "target": "mnopqrstuvwxyz123456"}})
	assert.Equal(t, dht.Dict{"t": "aa", "y": "r", "r": dht.Dict{"id": "0123456789abcdefghij", "nodes": "def456..."}}, resp)
}
