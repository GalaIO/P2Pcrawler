package dht

import (
	"github.com/GalaIO/P2Pcrawler/dht/krpc"
	"github.com/GalaIO/P2Pcrawler/misc"
	"net"
)

// response
func WithPingResponse(txId, nodeId string) krpc.Response {
	return krpc.WithResponse(txId, misc.Dict{"id": nodeId})
}

func WithFindNodeResponse(txId string, nodeId string, nodes []*NodeInfo) krpc.Response {
	return krpc.WithResponse(txId, misc.Dict{"id": nodeId, "nodes": joinNodeInfos(nodes)})
}

func WithGetPeersValsResponse(txId, nodeId, token string, addrs []*net.UDPAddr) krpc.Response {
	return krpc.WithResponse(txId, misc.Dict{"id": nodeId, "token": token, "values": joinPeerInfos(addrs)})
}

func WithGetPeersNodesResponse(txId, nodeId, token string, nodes []*NodeInfo) krpc.Response {
	return krpc.WithResponse(txId, misc.Dict{"id": nodeId, "token": token, "nodes": joinNodeInfos(nodes)})
}

func WithAnnouncePeerResponse(txId string, nodeId string) krpc.Response {
	return krpc.WithResponse(txId, misc.Dict{"id": nodeId})
}

// define query msg
func WithPingMsg(txId string, nodeId string, handler krpc.RpcHandlerFunc) krpc.Request {
	return krpc.NewBaseRequest(txId, "ping", misc.Dict{
		"id": nodeId,
	}, handler)
}

func WithFindNodeMsg(txId string, nodeId, target string, handler krpc.RpcHandlerFunc) krpc.Request {
	return krpc.NewBaseRequest(txId, "find_node", misc.Dict{
		"id":     nodeId,
		"target": target,
	}, handler)
}

func WithGetPeersMsg(txId, nodeId, infoHash string, handler krpc.RpcHandlerFunc) krpc.Request {
	return krpc.NewBaseRequest(txId, "get_peers", misc.Dict{
		"id":        nodeId,
		"info_hash": infoHash,
	}, handler)
}

func WithAnnouncePeerMsg(txId, nodeId, infoHash, token string, port int, handler krpc.RpcHandlerFunc) krpc.Request {
	return krpc.NewBaseRequest(txId, "announce_peer", misc.Dict{
		"id":           nodeId,
		"info_hash":    infoHash,
		"port":         port,
		"token":        token,
		"implied_port": 0,
	}, handler)
}
