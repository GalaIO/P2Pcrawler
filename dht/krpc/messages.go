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

type Request interface {
	TxId() string
	Type() string
	Body() misc.Dict
	Handler() RpcHandlerFunc
	RawData() misc.Dict
	fmt.Stringer
}

type BaseRequest struct {
	txId    string
	qType   string
	body    misc.Dict
	handler RpcHandlerFunc
}

func NewBaseRequest(txId, qtype string, body misc.Dict, handlerFunc RpcHandlerFunc) Request {
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

func (b *BaseRequest) Handler() RpcHandlerFunc {
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
	return NewBaseResponse(txId, false, resp)
}

// error
func WithErr(txId string, code int, errMsg string) Response {
	return NewBaseResponse(txId, true, misc.Dict{"code": code, "msg": errMsg})
}

func WithParamErr(txId string, msg string) Response {
	if len(msg) <= 0 {
		msg = "invalid param"
	}
	return WithErr(txId, int(misc.ProtocolErr), msg)
}
