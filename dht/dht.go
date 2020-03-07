package dht

import "github.com/GalaIO/P2Pcrawler/dht/krpc"

func Run() {

	krpc.RegisteHandler("ping", func(req krpc.Request) krpc.Response {
		return krpc.WithPingResponse(req.TxId(), krpc.LocalNodeId)
	})

	krpc.RegisteHandler("find_node", func(req krpc.Request) krpc.Response {
		return krpc.WithFindNodeResponse(req.TxId(), krpc.LocalNodeId, nil)
	})

	krpc.RegisteHandler("get_peers", func(req krpc.Request) krpc.Response {
		return krpc.WithGetPeersNodesResponse(req.TxId(), krpc.LocalNodeId, req.TxId(), nil)
	})

	krpc.RegisteHandler("announce_peer", func(req krpc.Request) krpc.Response {
		return krpc.WithAnnouncePeerResponse(req.TxId(), krpc.LocalNodeId)
	})

	krpc.BootStrap("87.98.162.88:6881", findNodeHandler)
}
