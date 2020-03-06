package krpc

import (
	"encoding/json"
	"fmt"
	"github.com/GalaIO/P2Pcrawler/misc"
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
	qType string
	body  misc.Dict
	isErr bool
}

func NewBaseResponse(txId, qType string, isErr bool, body misc.Dict) Response {
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

func (b *BaseResponse) Body() misc.Dict {
	return b.body
}

func (b *BaseResponse) Error() bool {
	return b.isErr
}

func (b *BaseResponse) RawData() misc.Dict {
	respType := "r"
	if b.isErr {
		respType = "e"
	}
	return misc.Dict{
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
func withResponse(txId string, resp misc.Dict) misc.Dict {
	return misc.Dict{
		"t": txId,
		"y": "r",
		"r": resp,
	}
}

// error
func withParamErr(msg string) misc.Dict {
	if len(msg) <= 0 {
		msg = "invalid param"
	}
	return withErr(misc.ProtocolErr, msg)
}

func withErr(code misc.ErrCode, errMsg string) misc.Dict {
	return misc.Dict{"t": "aa", "y": "e", "e": misc.List{code, errMsg}}
}

// define query msg
func withPingMsg(txId string, nodeId string, handler RespHandlerFunc) Request {
	return NewBaseRequest(txId, "ping", misc.Dict{
		"id": nodeId,
	}, handler)
}

func withFindNodeMsg(txId string, nodeId, target string, handler RespHandlerFunc) Request {
	return NewBaseRequest(txId, "find_node", misc.Dict{
		"id":     nodeId,
		"target": target,
	}, handler)
}
