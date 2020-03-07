package krpc

import (
	"encoding/json"
	"fmt"
	"github.com/GalaIO/P2Pcrawler/misc"
	"net"
)

type Response interface {
	TxId() string
	NodeId() string
	Body() misc.Dict
	Error() bool
	RawData() misc.Dict
	fmt.Stringer
}

type BaseResponse struct {
	txId  string
	body  misc.Dict
	isErr bool
}

func NewBaseResponse(txId string, isErr bool, body misc.Dict) Response {
	return &BaseResponse{
		txId:  txId,
		body:  body,
		isErr: isErr,
	}
}

func (b *BaseResponse) TxId() string {
	return b.txId
}

func (b *BaseResponse) NodeId() string {
	return b.body.GetString("id")
}

func (b *BaseResponse) Body() misc.Dict {
	return b.body
}

func (b *BaseResponse) Error() bool {
	return b.isErr
}

func (b *BaseResponse) RawData() misc.Dict {
	if !b.isErr {
		return misc.Dict{
			"t": b.txId,
			"y": "r",
			"r": b.body,
		}
	}
	code := b.body.GetInteger("code")
	msg := b.body.GetString("msg")
	return misc.Dict{
		"t": b.txId,
		"y": "e",
		"e": misc.List{code, msg},
	}
}

func (b *BaseResponse) String() string {
	bytes, err := json.Marshal(b.RawData())
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}

type RespHandlerFunc func(req Request, resp Response)

type Request interface {
	TxId() string
	Type() string
	Body() misc.Dict
	Handler() RespHandlerFunc
	RawData() misc.Dict
	fmt.Stringer
}

type BaseRequest struct {
	txId    string
	qType   string
	body    misc.Dict
	handler RespHandlerFunc
}

func NewBaseRequest(txId, qtype string, body misc.Dict, handlerFunc RespHandlerFunc) Request {
	return &BaseRequest{
		txId:    txId,
		qType:   qtype,
		body:    body,
		handler: handlerFunc,
	}
}

func (b *BaseRequest) TxId() string {
	return b.txId
}

func (b *BaseRequest) Type() string {
	return b.qType
}

func (b *BaseRequest) Body() misc.Dict {
	return b.body
}

func (b *BaseRequest) Handler() RespHandlerFunc {
	return b.handler
}

func (b *BaseRequest) RawData() misc.Dict {
	return misc.Dict{
		"t": b.txId,
		"y": "q",
		"q": b.qType,
		"a": b.body,
	}
}

func (b *BaseRequest) String() string {
	bytes, err := json.Marshal(b)
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}

// response
func WithResponse(txId string, resp misc.Dict) Response {
	return &BaseResponse{
		txId:  txId,
		body:  resp,
		isErr: false,
	}
}
func WithPingResponse(txId, nodeId string) Response {
	return &BaseResponse{
		txId:  txId,
		body:  misc.Dict{"id": nodeId},
		isErr: false,
	}
}

func WithFindNodeResponse(txId string, nodeId string, nodes []*NodeInfo) Response {
	return &BaseResponse{
		txId:  txId,
		body:  misc.Dict{"id": nodeId, "nodes": joinNodeInfos(nodes)},
		isErr: false,
	}
}

func WithGetPeersValsResponse(txId, nodeId, token string, addrs []*net.UDPAddr) Response {
	return &BaseResponse{
		txId:  txId,
		body:  misc.Dict{"id": nodeId, "token": token, "values": joinPeerInfos(addrs)},
		isErr: false,
	}
}

func WithGetPeersNodesResponse(txId, nodeId, token string, nodes []*NodeInfo) Response {
	return &BaseResponse{
		txId:  txId,
		body:  misc.Dict{"id": nodeId, "token": token, "nodes": joinNodeInfos(nodes)},
		isErr: false,
	}
}

func WithAnnouncePeerResponse(txId string, nodeId string) Response {
	return &BaseResponse{
		txId:  txId,
		body:  misc.Dict{"id": nodeId},
		isErr: false,
	}
}

// error
func WithErr(txId string, code int, errMsg string) Response {
	return &BaseResponse{
		txId:  txId,
		body:  misc.Dict{"code": code, "msg": errMsg},
		isErr: true,
	}
}

func WithParamErr(txId string, msg string) Response {
	if len(msg) <= 0 {
		msg = "invalid param"
	}
	return WithErr(txId, int(misc.ProtocolErr), msg)
}

// define query msg
func WithPingMsg(txId string, nodeId string, handler RespHandlerFunc) Request {
	return NewBaseRequest(txId, "ping", misc.Dict{
		"id": nodeId,
	}, handler)
}

func WithFindNodeMsg(txId string, nodeId, target string, handler RespHandlerFunc) Request {
	return NewBaseRequest(txId, "find_node", misc.Dict{
		"id":     nodeId,
		"target": target,
	}, handler)
}

func WithGetPeersMsg(txId, nodeId, infoHash string, handler RespHandlerFunc) Request {
	return NewBaseRequest(txId, "get_peers", misc.Dict{
		"id":        nodeId,
		"info_hash": infoHash,
	}, handler)
}

func WithAnnouncePeerMsg(txId, nodeId, infoHash, token string, port int, handler RespHandlerFunc) Request {
	return NewBaseRequest(txId, "announce_peer", misc.Dict{
		"id":           nodeId,
		"info_hash":    infoHash,
		"port":         port,
		"token":        token,
		"implied_port": 0,
	}, handler)
}
