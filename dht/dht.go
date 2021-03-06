package dht

import (
	"fmt"
	"github.com/GalaIO/P2Pcrawler/config"
	"github.com/GalaIO/P2Pcrawler/dht/krpc"
	"github.com/GalaIO/P2Pcrawler/misc"
	"net"
)

var dhtLogger = misc.GetLogger().SetPrefix("dht")
var localAddr = fmt.Sprintf(":%d", config.DhtConfig().Port)
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
	rpcServer.RegisteHandler("ping", pingHandler)
	rpcServer.RegisteHandler("find_node", findNodeHandler)
	rpcServer.RegisteHandler("get_peers", getPeersHandler)
	rpcServer.RegisteHandler("announce_peer", announcePeerHandler)

	dhtLogger.Info("enter fetchTorrent loop...", nil)
	go fetchTorrent()

	misc.RegisterShutDownClean(func() {
		rpcServer.Close()
		dhtLogger.Info("rpcserver closed...", nil)
	})
	dhtLogger.Info("enter listen loop...", nil)
	rpcServer.Listen()
}

// bootstrap find myself
func BootStrap(host string) {
	dhtLogger.Info("bootstrap dht...", misc.Dict{"host": host})
	msg := WithFindNodeMsg(txIdGenerator.Next(), localNodeId, localNodeId, genericRespHandler)
	raddr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		dhtLogger.Panic("bootstrap resolve host err", misc.Dict{"host": host, "err": err})
	}

	err = rpcServer.Query(raddr, msg)
	if err != nil {
		dhtLogger.Panic("bootstrap findnode err", misc.Dict{"host": host, "err": err})
	}
}
