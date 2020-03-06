package misc

import "strconv"

type ErrCode int

const (
	NoneErr ErrCode = 200 + iota
	GenericErr
	ServerErr
	ProtocolErr
	UnknowMethod
)

type Error struct {
	code       ErrCode
	msg        string
	context    List
	printStack bool
}

func (d *Error) Error() string {
	return strconv.Itoa(int(d.code)) + ":" + d.msg
}

func NewError(code ErrCode, msg string, printStack bool, params ...interface{}) *Error {
	return &Error{
		code:       code,
		msg:        msg,
		context:    params,
		printStack: printStack,
	}
}

func panicSysErr(errMsg string, params ...interface{}) {
	panic(NewError(GenericErr, errMsg, true, params))
}

func PanicBizErr(errMsg string, params ...interface{}) {
	panic(NewError(GenericErr, errMsg, false, params))
}
