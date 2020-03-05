package krpc

import (
	"encoding/json"
	"fmt"
	"github.com/GalaIO/P2Pcrawler/dht"
)

type Response interface {
	TxId() string
	NodeId() string
	Body() dht.Dict
	Error() bool
	RawData() dht.Dict
	fmt.Stringer
}

type BaseResponse struct {
	txId  string
	qType string
	body  dht.Dict
	isErr bool
}

func NewBaseResponse(txId, qType string, isErr bool, body dht.Dict) Response {
	return &BaseResponse{
		txId:  txId,
		qType: qType,
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

func (b *BaseResponse) Body() dht.Dict {
	return b.body
}

func (b *BaseResponse) Error() bool {
	return b.isErr
}

func (b *BaseResponse) RawData() dht.Dict {
	respType := "r"
	if b.isErr {
		respType = "e"
	}
	return dht.Dict{
		"t": b.txId,
		"y": respType,
		"r": b.body,
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
	Body() dht.Dict
	Handler() RespHandlerFunc
	RawData() dht.Dict
	fmt.Stringer
}

type BaseRequest struct {
	txId    string
	qType   string
	body    dht.Dict
	handler RespHandlerFunc
}

func NewBaseRequest(txId, qtype string, body dht.Dict, handlerFunc RespHandlerFunc) Request {
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

func (b *BaseRequest) Body() dht.Dict {
	return b.body
}

func (b *BaseRequest) Handler() RespHandlerFunc {
	return b.handler
}

func (b *BaseRequest) RawData() dht.Dict {
	return dht.Dict{
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
func withResponse(txId string, resp dht.Dict) dht.Dict {
	return dht.Dict{
		"t": txId,
		"y": "r",
		"r": resp,
	}
}

// error
func withParamErr(msg string) dht.Dict {
	if len(msg) <= 0 {
		msg = "invalid param"
	}
	return withErr(dht.ProtocolErr, msg)
}

func withErr(code dht.DhtErrCode, errMsg string) dht.Dict {
	return dht.Dict{"t": "aa", "y": "e", "e": dht.List{code, errMsg}}
}

// define query msg
func withPingMsg(txId string, nodeId string, handler RespHandlerFunc) Request {
	return NewBaseRequest(txId, "ping", dht.Dict{
		"id": nodeId,
	}, handler)
}

func withFindNodeMsg(txId string, nodeId, target string, handler RespHandlerFunc) Request {
	return NewBaseRequest(txId, "find_node", dht.Dict{
		"id":     nodeId,
		"target": target,
	}, handler)
}
