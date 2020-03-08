package dht

import (
	"encoding/hex"
	"github.com/GalaIO/P2Pcrawler/dht/krpc"
	"github.com/GalaIO/P2Pcrawler/misc"
	"net"
)

var txIdGen = NewTxIdGenerator(100)
var dhtLogger = misc.GetLogger().SetPrefix("dht")
var localAddr = ":21000"
var localNodeId = generateNodeId("galaio.p2pclawer")

const maxRouteTableSize int = 160
const maxBucketLen int = 8

var rtable = NewRouteTable(NewNodeInfoFromHost(localNodeId, localAddr), maxBucketLen, maxRouteTableSize)
var rpcServer = krpc.NewRpcServer(localAddr)
var txIdGenerator = NewTxIdGenerator(1000)

func Run() {
	dhtLogger.Info("start run dht...", nil)

	dhtLogger.Info("register middleware...", nil)
	//rpcServer.UseRespHandlerMiddleware(renderRespMiddleware)

	dhtLogger.Info("register request handler...", nil)
	registerRequestHandler()

	dhtLogger.Info("enter listen loop...", nil)
	rpcServer.Listen()
}

func registerRequestHandler() {
	rpcServer.RegisteHandler("ping", func(ctx *krpc.RpcContext) {
		req := ctx.Request()
		dhtLogger.Info("get ping request", misc.Dict{"txId": req.TxId()})
		ctx.WriteAs(WithPingResponse(req.TxId(), localNodeId))
	})
	rpcServer.RegisteHandler("find_node", func(ctx *krpc.RpcContext) {
		req := ctx.Request()
		dhtLogger.Info("get find_node request", misc.Dict{"txId": req.TxId()})
		ctx.WriteAs(WithFindNodeResponse(req.TxId(), localNodeId, nil))
	})
	rpcServer.RegisteHandler("get_peers", func(ctx *krpc.RpcContext) {
		req := ctx.Request()
		body := req.Body()
		infoHash := body.GetString("info_hash")
		hash := hex.EncodeToString([]byte(infoHash))
		dhtLogger.Info("get get_peers request", misc.Dict{"txId": req.TxId(), "infoHash": hash})
		ctx.WriteAs(WithGetPeersNodesResponse(req.TxId(), localNodeId, req.TxId(), nil))
	})
	rpcServer.RegisteHandler("announce_peer", func(ctx *krpc.RpcContext) {
		req := ctx.Request()
		body := req.Body()
		infoHash := body.GetString("info_hash")
		hash := hex.EncodeToString([]byte(infoHash))
		dhtLogger.Info("get announce_peer request", misc.Dict{"txId": req.TxId(), "infoHash": hash, "port": body.GetInteger("port")})
		ctx.WriteAs(WithAnnouncePeerResponse(req.TxId(), localNodeId))
	})
}

// bootstrap find myself
func BootStrap(host string) {
	dhtLogger.Info("bootstrap dht...", nil)
	msg := WithFindNodeMsg(txIdGen.Next(), localNodeId, localNodeId, genericRespHandler)
	raddr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		dhtLogger.Panic("bootstrap resolve host err", misc.Dict{"host": host, "err": err})
	}

	err = rpcServer.Query(raddr, msg)
	if err != nil {
		dhtLogger.Panic("bootstrap findnode err", misc.Dict{"host": host, "err": err})
	}
}
