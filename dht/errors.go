package dht

import "strconv"

type DhtErrCode int

const (
	NoneErr DhtErrCode = 200 + iota
	GenericErr
	ServerErr
	ProtocolErr
	UnknowMethod
)

type DhtError struct {
	code       DhtErrCode
	msg        string
	context    List
	printStack bool
}

func (d *DhtError) Error() string {
	return strconv.Itoa(int(d.code)) + ":" + d.msg
}

func NewDhtErr(code DhtErrCode, msg string, printStack bool, params ...interface{}) *DhtError {
	return &DhtError{
		code:       code,
		msg:        msg,
		context:    params,
		printStack: printStack,
	}
}

func panicSysErr(errMsg string, params ...interface{}) {
	panic(NewDhtErr(GenericErr, errMsg, true, params))
}

func panicBizErr(errMsg string, params ...interface{}) {
	panic(NewDhtErr(GenericErr, errMsg, false, params))
}

func withParamErr(msg string) Dict {
	if len(msg) <= 0 {
		msg = "invalid param"
	}
	return withErr(ProtocolErr, msg)
}

func withErr(code DhtErrCode, errMsg string) Dict {
	return Dict{"t": "aa", "y": "e", "e": List{code, errMsg}}
}
