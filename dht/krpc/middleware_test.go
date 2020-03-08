package krpc

import (
	"github.com/GalaIO/P2Pcrawler/misc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReqHandlerMiddleware(t *testing.T) {
	execute := ""
	rpcServer := NewRpcServer(":20001")
	defer rpcServer.Close()
	rpcServer.UseReqHandlerMiddleware(func(ctx *RpcContext) {
		execute += "a"
		ctx.Next()
		execute += "a"
	}, func(ctx *RpcContext) {
		execute += "b"
		ctx.Next()
		execute += "b"
	})

	request := NewBaseRequest("aa", "ping", misc.Dict{
		"id": "test",
	}, nil)
	resp := rpcServer.doRequestHandle(request, func(ctx *RpcContext) {
		execute += "main"
		ctx.WriteAs(WithResponse("test", misc.Dict{}))
	})

	assert.Equal(t, "abmainba", execute)
	assert.Equal(t, WithResponse("test", misc.Dict{}), resp)
}

func TestReqHandlerMiddlewareWithNone(t *testing.T) {
	execute := ""
	rpcServer := NewRpcServer(":20001")
	defer rpcServer.Close()

	request := NewBaseRequest("aa", "ping", misc.Dict{
		"id": "test",
	}, nil)
	resp := rpcServer.doRequestHandle(request, func(ctx *RpcContext) {
		execute += "main"
		ctx.WriteAs(WithResponse("test", misc.Dict{}))
	})

	assert.Equal(t, "main", execute)
	assert.Equal(t, WithResponse("test", misc.Dict{}), resp)
}

func TestRespHandlerMiddleware(t *testing.T) {
	execute := ""
	rpcServer := NewRpcServer(":20001")
	defer rpcServer.Close()
	rpcServer.UseRespHandlerMiddleware(func(ctx *RpcContext) {
		execute += "a"
		ctx.Next()
		execute += "a"
	}, func(ctx *RpcContext) {
		execute += "b"
		ctx.Next()
		execute += "b"
	})

	var resp Response
	request := NewBaseRequest("aa", "ping", misc.Dict{
		"id": "test",
	}, func(ctx *RpcContext) {
		execute += "main"
		resp = ctx.Response()
	})
	rpcServer.doResponseHandle(request, WithResponse("aa", misc.Dict{}))

	assert.Equal(t, "abmainba", execute)
	assert.Equal(t, WithResponse("aa", misc.Dict{}), resp)
}

func TestRespHandlerMiddlewareWithNone(t *testing.T) {
	execute := ""
	rpcServer := NewRpcServer(":20001")
	defer rpcServer.Close()

	var resp Response
	request := NewBaseRequest("aa", "ping", misc.Dict{
		"id": "test",
	}, func(ctx *RpcContext) {
		execute += "main"
		resp = ctx.Response()
	})
	rpcServer.doResponseHandle(request, WithResponse("aa", misc.Dict{}))

	assert.Equal(t, "main", execute)
	assert.Equal(t, WithResponse("aa", misc.Dict{}), resp)
}
