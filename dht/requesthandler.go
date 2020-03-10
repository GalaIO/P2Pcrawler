package dht

import (
	"encoding/hex"
	"github.com/GalaIO/P2Pcrawler/dht/krpc"
	"github.com/GalaIO/P2Pcrawler/misc"
)

var infoHashLogger = misc.NewZapLogger("infohash", "info")

func pingHandler(ctx *krpc.RpcContext) {
	req := ctx.Request()
	dhtLogger.Info("get ping request", misc.Dict{"txId": req.TxId()})
	ctx.WriteAs(WithPingResponse(req.TxId(), localNodeId))
}

func findNodeHandler(ctx *krpc.RpcContext) {
	req := ctx.Request()
	dhtLogger.Info("get find_node request", misc.Dict{"txId": req.TxId()})
	ctx.WriteAs(WithFindNodeResponse(req.TxId(), localNodeId, nil))
}

func getPeersHandler(ctx *krpc.RpcContext) {
	req := ctx.Request()
	body := req.Body()
	infoHash := body.GetString("info_hash")
	hash := hex.EncodeToString([]byte(infoHash))
	dhtLogger.Info("get get_peers request", misc.Dict{"txId": req.TxId(), "infoHash": hash})
	infoHashLogger.Info("get_peer", misc.Dict{"addr": ctx.RemoteAddr().String(), "infoHash": hash})
	ctx.WriteAs(WithGetPeersNodesResponse(req.TxId(), localNodeId, req.TxId(), nil))
}

func announcePeerHandler(ctx *krpc.RpcContext) {
	req := ctx.Request()
	body := req.Body()
	infoHash := body.GetString("info_hash")
	hash := hex.EncodeToString([]byte(infoHash))
	port := body.GetInteger("port")
	dhtLogger.Info("get announce_peer request", misc.Dict{"txId": req.TxId(), "infoHash": hash, "port": port})
	infoHashLogger.Info("announce_peer", misc.Dict{"addr": ctx.RemoteAddr().String(), "port": port, "infoHash": hash})
	ctx.WriteAs(WithAnnouncePeerResponse(req.TxId(), localNodeId))
	// to chan
	recvInfoHash <- ctx
}
