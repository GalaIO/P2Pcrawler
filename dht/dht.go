package dht

import (
	"github.com/GalaIO/P2Pcrawler/dht/krpc"
	"github.com/GalaIO/P2Pcrawler/misc"
	"net"
)

var txIdGen = NewTxIdGenerator(100)
var dhtLogger = misc.GetLogger().SetPrefix("dht")

func Run() {

	dhtLogger.Info("start run dht", nil)
	krpc.RegisteHandler("ping", func(req krpc.Request) krpc.Response {
		return WithPingResponse(req.TxId(), LocalNodeId)
	})

	krpc.RegisteHandler("find_node", func(req krpc.Request) krpc.Response {
		return WithFindNodeResponse(req.TxId(), LocalNodeId, nil)
	})

	krpc.RegisteHandler("get_peers", func(req krpc.Request) krpc.Response {
		return WithGetPeersNodesResponse(req.TxId(), LocalNodeId, req.TxId(), nil)
	})

	krpc.RegisteHandler("announce_peer", func(req krpc.Request) krpc.Response {
		return WithAnnouncePeerResponse(req.TxId(), LocalNodeId)
	})

	BootStrap("87.98.162.88:6881", findNodeHandler)
}

// bootstrap find myself
func BootStrap(host string, handler krpc.RespHandlerFunc) {
	msg := WithFindNodeMsg(txIdGen.Next(), LocalNodeId, LocalNodeId, handler)
	raddr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		dhtLogger.Panic("bootstrap resolve host err", misc.Dict{"host": host, "err": err})
	}

	err = krpc.Query(raddr, msg)
	if err != nil {
		dhtLogger.Panic("bootstrap findnode err", misc.Dict{"host": host, "err": err})
	}
}
