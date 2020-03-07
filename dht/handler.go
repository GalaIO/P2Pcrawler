package dht

import "github.com/GalaIO/P2Pcrawler/dht/krpc"

// reconstruct response
func renderResp(req krpc.Request, resp krpc.Response) krpc.Response {
	txId := resp.TxId()
	nodeId := resp.NodeId()
	body := resp.Body()
	switch req.Type() {
	case "ping":
		return WithPingResponse(txId, nodeId)
	case "find_node":
		nodes := body.GetString("nodes")
		return WithFindNodeResponse(txId, nodeId, parseNodeInfo(nodes))
	case "get_peers":
		existVals := body.Exist("values")
		if existVals {
			vals := body.GetList("values")
			return WithGetPeersValsResponse(txId, nodeId, txId, parsePeerInfo(vals))
		}
		nodes := body.GetString("nodes")
		return WithGetPeersNodesResponse(txId, nodeId, txId, parseNodeInfo(nodes))
	case "announce_peer":
		return WithAnnouncePeerResponse(txId, nodeId)
	}
	return resp
}

func findNodeHandler(req krpc.Request, resp krpc.Response) {

}
