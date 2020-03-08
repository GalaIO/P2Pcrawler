package krpc

import (
	"github.com/GalaIO/P2Pcrawler/misc"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testNodeId = "abcdefghij0123456789"

func TestHandlePingMsg(t *testing.T) {
	rpcServer := NewRpcServer(":21000")
	defer rpcServer.Close()
	rpcServer.RegisteHandler("ping", func(ctx *RpcContext) {
		req := ctx.Request()
		ctx.WriteAs(WithResponse(req.TxId(), misc.Dict{"id": "mnopqrstuvwxyz123456"}))
	})

	resp := rpcServer.requestHandle(nil, misc.Dict{"t": "aa", "y": "q", "q": "ping", "a": misc.Dict{"id": "abcdefghij0123456789"}})
	assert.Equal(t, misc.Dict{"t": "aa", "y": "r", "r": misc.Dict{"id": "mnopqrstuvwxyz123456"}}, resp.RawData())

	resp = rpcServer.requestHandle(nil, misc.Dict{"t": "aa", "y": "q", "q": "ping"})
	assert.Equal(t, misc.Dict{"t": "aa", "y": "e", "e": misc.List{203, "201:cannot find a's val"}}, resp.RawData())
}

func TestHandleFindNodeMsg(t *testing.T) {
	rpcServer := NewRpcServer(":21000")
	defer rpcServer.Close()
	rpcServer.RegisteHandler("find_node", func(ctx *RpcContext) {
		req := ctx.Request()
		ctx.WriteAs(WithResponse(req.TxId(), misc.Dict{"id": "0123456789abcdefghij", "nodes": "def456..."}))
	})

	resp := rpcServer.requestHandle(nil, misc.Dict{"t": "aa", "y": "q", "q": "find_node", "a": misc.Dict{"id": "abcdefghij0123456789", "target": "mnopqrstuvwxyz123456"}})
	assert.Equal(t, misc.Dict{"t": "aa", "y": "r", "r": misc.Dict{"id": "0123456789abcdefghij", "nodes": "def456..."}}, resp.RawData())
}
