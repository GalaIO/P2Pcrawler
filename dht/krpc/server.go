package krpc

import (
	"github.com/GalaIO/P2Pcrawler/misc"
	"net"
)

var serverLogger = misc.GetLogger().SetPrefix("server")
var supportQueryType = misc.List{"ping", "find_node", "get_peers", "announce_peer"}

var serverConn *UdpServer
var reqestHandlerRouter = misc.NewSyncMap(4)
var requestMapping = misc.NewSyncMap(100)

type ReqHandlerFunc func(req Request) Response

func RegisteHandler(qType string, handler ReqHandlerFunc) {
	if !supportQueryType.ContainsString(qType) {
		panic("cannot support the query type")
	}

	if handler == nil {
		panic("register fail, handler is nil")
	}

	reqestHandlerRouter.Put(qType, handler)
}

// the server main routinue, will listen and handle msg
func Server() {
	serverConn = StartUp(":21000")
	for {
		packet := <-serverConn.RecvChan()
		serverLogger.Info("<<<  Bytes received", misc.Dict{"from": packet.Addr.String(), "len": len(packet.Bytes)})

		go recvPacketHandle(packet)
	}
}

func Query(raddr *net.UDPAddr, req Request) error {
	requestMapping.Put(req.TxId(), req)
	raw, err := misc.EncodeDict(req.RawData())
	if err != nil {
		serverLogger.Error("encode request err", misc.Dict{"to": raddr.String(), "query": req.String(), "err": err})
		return err
	}
	err = serverConn.SendPacket([]byte(raw), raddr)
	if err != nil {
		serverLogger.Error("send query packet err", misc.Dict{"to": raddr.String(), "query": req.String(), "err": err})
		return err
	}
	serverLogger.Info(">>>  Bytes sended", misc.Dict{"to": raddr.String(), "len": len(raw)})
	return nil
}

func recvPacketHandle(packet RecvPacket) {

	defer func() {
		if err := recover(); err != nil {
			serverLogger.Error("recv packet handle panic", misc.Dict{"from": packet.Addr.String(), "err": err})
		}
	}()

	// parse packet
	dict, err := misc.DecodeDict(string(packet.Bytes))
	if err != nil {
		serverLogger.Error("decode bencode err", misc.Dict{"from": packet.Addr.String(), "err": err})
		return
	}
	if exist := dict.Exist("y"); !exist {
		serverLogger.Error("cannot handle packet err", misc.Dict{"from": packet.Addr.String()})
		return
	}
	switch dict.GetString("y") {
	case "q":
		ret := requestHandle(dict)
		bytes, err := misc.EncodeDict(ret.RawData())
		if err != nil {
			serverLogger.Error("encode response err", misc.Dict{"from": packet.Addr.String(), "dict": ret.String()})
		}
		err = serverConn.SendPacket([]byte(bytes), packet.Addr)
		serverLogger.Info(">>>  Bytes sended", misc.Dict{"to": packet.Addr.String(), "len": len(bytes)})
	case "r":
		responseHandle(dict)
	case "e":
		responseErrHandle(dict)
	}
}

// err handler
func responseErrHandle(err misc.Dict) {
	// parse header
	txId := err.GetString("t")
	list := err.GetList("e")
	serverLogger.Error("return err", misc.Dict{"txId": txId, "err": list})
}

// response handler
func responseHandle(dict misc.Dict) {

	// parse header
	txId := dict.GetString("t")
	body := dict.GetDict("r")
	nodeId := body.GetString("nodeId")
	handler, exist := requestMapping.Get(txId)

	if !exist {
		serverLogger.Error("cannot match request", misc.Dict{"txId": txId, "nodeId": nodeId})
		return
	}

	req := handler.(Request)
	resp := WithResponse(txId, body)
	req.Handler()(req, resp)
}

// query handler
func requestHandle(resp misc.Dict) (ret Response) {

	txId := resp.GetString("t")
	defer func() {
		if err := recover(); err != nil {
			serverLogger.Error("request handle panic", misc.Dict{"err": err})
			dhtError, ok := err.(*misc.Error)
			if !ok {
				panic(err)
			}
			ret = WithParamErr(txId, dhtError.Error())
		}
	}()

	// parse header
	queryType := resp.GetString("q")
	if !supportQueryType.ContainsString(queryType) {
		return WithParamErr(txId, "donnot support <"+queryType+"> query type")
	}
	body := resp.GetDict("a")
	sourceId := body.GetString("id")
	if len(sourceId) != 20 {
		return WithParamErr(txId, "id format err")
	}

	req := NewBaseRequest(txId, queryType, body, nil)
	handler, exist := reqestHandlerRouter.Get(queryType)
	if !exist {
		return WithParamErr(txId, "cannot handle not match handler")
	}
	handlerFunc := handler.(ReqHandlerFunc)
	return handlerFunc(req)
}
