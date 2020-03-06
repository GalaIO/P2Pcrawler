package krpc

import (
	"github.com/GalaIO/P2Pcrawler/misc"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testNodeId = "abcdefghij0123456789"

func TestHandlePingMsg(t *testing.T) {
	RegisteQueryHandler("ping", func(ctx *QueryCtx) misc.Dict {
		return misc.Dict{"id": "mnopqrstuvwxyz123456"}
	})

	resp := queriesHandle(misc.Dict{"t": "aa", "y": "q", "q": "ping", "a": misc.Dict{"id": "abcdefghij0123456789"}})
	assert.Equal(t, misc.Dict{"t": "aa", "y": "r", "r": misc.Dict{"id": "mnopqrstuvwxyz123456"}}, resp)

	resp = queriesHandle(misc.Dict{"t": "aa", "y": "q", "q": "ping"})
	assert.Equal(t, misc.Dict{"t": "aa", "y": "e", "e": misc.List{misc.ProtocolErr, "201:cannot find a's val"}}, resp)
}

func TestHandleFindNodeMsg(t *testing.T) {
	RegisteQueryHandler("find_node", func(ctx *QueryCtx) misc.Dict {
		return misc.Dict{"id": "0123456789abcdefghij", "nodes": "def456..."}
	})

	resp := queriesHandle(misc.Dict{"t": "aa", "y": "q", "q": "find_node", "a": misc.Dict{"id": "abcdefghij0123456789", "target": "mnopqrstuvwxyz123456"}})
	assert.Equal(t, misc.Dict{"t": "aa", "y": "r", "r": misc.Dict{"id": "0123456789abcdefghij", "nodes": "def456..."}}, resp)
}
