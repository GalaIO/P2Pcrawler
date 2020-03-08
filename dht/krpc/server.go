package krpc

import (
	"github.com/GalaIO/P2Pcrawler/misc"
	"net"
)

var serverLogger = misc.GetLogger().SetPrefix("server")
var supportQueryType = misc.List{"ping", "find_node", "get_peers", "announce_peer"}

type RpcHandlerFunc func(ctx *RpcContext)
type RpcServer struct {
	udpConn             *UdpServer
	reqestHandlerRouter *misc.SyncMap
	requestMapping      *misc.SyncMap
	reqHandlerChain     []RpcHandlerFunc
	respHandlerChain    []RpcHandlerFunc
}

func NewRpcServer(laddr string) *RpcServer {
	udpConn := StartUdpServer(laddr)
	return &RpcServer{
		udpConn:             udpConn,
		reqestHandlerRouter: misc.NewSyncMap(4),
		requestMapping:      misc.NewSyncMap(100),
		reqHandlerChain:     make([]RpcHandlerFunc, 0, 16),
		respHandlerChain:    make([]RpcHandlerFunc, 0, 16),
	}
}

func (s *RpcServer) Listen() {
	for {
		packet := <-s.udpConn.RecvChan()
		serverLogger.Info("<<<  Bytes received", misc.Dict{"from": packet.Addr.String(), "len": len(packet.Bytes)})

		go s.recvPacketHandle(packet)
	}
}

func (s *RpcServer) Close() {
	if s.udpConn != nil {
		s.udpConn.Close()
		serverLogger.Info("close udp conn", nil)
	}
}

func (s *RpcServer) UseReqHandlerMiddleware(handlerChain ...RpcHandlerFunc) {
	s.reqHandlerChain = append(s.reqHandlerChain, handlerChain...)
}

func (s *RpcServer) UseRespHandlerMiddleware(handlerChain ...RpcHandlerFunc) {
	s.respHandlerChain = append(s.respHandlerChain, handlerChain...)
}

func (s *RpcServer) doRequestHandle(req Request, reqHandler RpcHandlerFunc) Response {
	executeChain := make([]RpcHandlerFunc, 0, len(s.reqHandlerChain))
	executeChain = append(executeChain, s.reqHandlerChain...)
	executeChain = append(executeChain, reqHandler)
	ctx := NewReqContext(executeChain, req, nil)
	ctx.Next()
	return ctx.resp
}

func (s *RpcServer) doResponseHandle(req Request, resp Response) {
	executeChain := make([]RpcHandlerFunc, 0, len(s.reqHandlerChain))
	executeChain = append(executeChain, s.respHandlerChain...)
	executeChain = append(executeChain, req.Handler())
	ctx := NewReqContext(executeChain, req, resp)
	ctx.Next()
}

func (s *RpcServer) RegisteHandler(qType string, handler RpcHandlerFunc) {
	if !supportQueryType.ContainsString(qType) {
		panic("cannot support the query type")
	}

	if handler == nil {
		panic("register fail, handler is nil")
	}

	s.reqestHandlerRouter.Put(qType, handler)
}

// send query msg
func (s *RpcServer) Query(raddr *net.UDPAddr, req Request) error {
	s.requestMapping.Put(req.TxId(), req)
	raw, err := misc.EncodeDict(req.RawData())
	if err != nil {
		serverLogger.Error("encode request err", misc.Dict{"to": raddr.String(), "query": req.String(), "err": err})
		return err
	}
	err = s.udpConn.SendPacket([]byte(raw), raddr)
	if err != nil {
		serverLogger.Error("send query packet err", misc.Dict{"to": raddr.String(), "query": req.String(), "err": err})
		return err
	}
	serverLogger.Info(">>>  Bytes sended", misc.Dict{"to": raddr.String(), "len": len(raw)})
	return nil
}

func (s *RpcServer) recvPacketHandle(packet *RecvPacket) {

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
		ret := s.requestHandle(dict)
		bytes, err := misc.EncodeDict(ret.RawData())
		if err != nil {
			serverLogger.Error("encode response err", misc.Dict{"from": packet.Addr.String(), "dict": ret.String()})
		}
		err = s.udpConn.SendPacket([]byte(bytes), packet.Addr)
		serverLogger.Info(">>>  Bytes sended", misc.Dict{"to": packet.Addr.String(), "len": len(bytes)})
	case "r":
		s.responseHandle(dict)
	case "e":
		s.responseErrHandle(dict)
	}
}

// err handler
func (s *RpcServer) responseErrHandle(err misc.Dict) {
	// parse header
	txId := err.GetString("t")
	list := err.GetList("e")
	serverLogger.Error("return err", misc.Dict{"txId": txId, "err": list})
}

// response handler
func (s *RpcServer) responseHandle(dict misc.Dict) {

	// parse header
	txId := dict.GetString("t")
	body := dict.GetDict("r")
	nodeId := body.GetString("nodeId")
	handler, exist := s.requestMapping.Get(txId)

	if !exist {
		serverLogger.Error("cannot match request", misc.Dict{"txId": txId, "nodeId": nodeId})
		return
	}

	req := handler.(Request)
	resp := WithResponse(txId, body)
	s.doResponseHandle(req, resp)
}

// query handler
func (s *RpcServer) requestHandle(resp misc.Dict) (ret Response) {

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
	handler, exist := s.reqestHandlerRouter.Get(queryType)
	if !exist {
		return WithParamErr(txId, "cannot handle not match handler")
	}
	handlerFunc := handler.(RpcHandlerFunc)
	return s.doRequestHandle(req, handlerFunc)
}
