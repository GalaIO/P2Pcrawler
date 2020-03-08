package dht

import (
	"github.com/GalaIO/P2Pcrawler/dht/krpc"
	"github.com/GalaIO/P2Pcrawler/misc"
	"net"
)

func genericRespHandler(ctx *krpc.RpcContext) {
	req := ctx.Request()
	resp := ctx.Response()

	txId := resp.TxId()
	nodeId := resp.NodeId()
	body := resp.Body()
	dhtLogger.Info("handle response", misc.Dict{"txid": txId, "queryType": req.Type()})
	switch req.Type() {
	case "ping":
		pingRespHandler(ctx, nodeId)
	case "find_node":
		nodes := body.GetString("nodes")
		findNodeRespHandler(ctx, parseNodeInfo(nodes))
	case "get_peers":
		existVals := body.Exist("values")
		if existVals {
			vals := body.GetList("values")
			getPeerValRespHandler(ctx, parsePeerInfo(vals))
			return
		}
		nodes := body.GetString("nodes")
		getPeerNodesRespHandler(ctx, parseNodeInfo(nodes))
	case "announce_peer":
		announcePeerRespHandler(ctx, nodeId)
	}
	dhtLogger.Info("cannot match request ignore", nil)
}

func findNodeRespHandler(ctx *krpc.RpcContext, nodes []*NodeInfo) {
	for _, n := range nodes {
		err := rtable.AddNode(n)
		if err != nil {
			dhtLogger.Info("add node err", misc.Dict{"err": err, "tableSize": rtable.Size()})
			continue
		}
		err = rpcServer.Query(n.Addr, WithFindNodeMsg(txIdGenerator.Next(), localNodeId, localNodeId, genericRespHandler))
		if err != nil {
			dhtLogger.Error("recurve find node err", misc.Dict{"err": err, "tableSize": rtable.Size()})
		}
	}
	dhtLogger.Info("add nodes", misc.Dict{"size": len(nodes), "tableSize": rtable.Size()})
}

func pingRespHandler(ctx *krpc.RpcContext, nodeId string) {
	dhtLogger.Info("ping response", misc.Dict{"nodeId": nodeId})
}

func getPeerValRespHandler(ctx *krpc.RpcContext, vals []*net.UDPAddr) {
	dhtLogger.Info("getpeer with peers has not implement", misc.Dict{"values": vals})
}

func getPeerNodesRespHandler(ctx *krpc.RpcContext, nodes []*NodeInfo) {
	dhtLogger.Info("getpeer with nodes has not implement", misc.Dict{"nodes": nodes})
}

func announcePeerRespHandler(ctx *krpc.RpcContext, nodeId string) {
	dhtLogger.Info("announce peer response", misc.Dict{"nodeId": nodeId})
}
