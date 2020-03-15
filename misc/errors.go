package misc

import (
	"runtime/debug"
	"strconv"
)

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
	errString := strconv.Itoa(int(d.code)) + ": " + d.msg
	if d.context != nil {
		errString += ", params: "
		for _, val := range d.context {
			errString += ToString(val) + "|"
		}
	}
	errString += "\r\nstacks:\r\n" + Bytes2Str(debug.Stack())
	return errString
}

func NewError(code ErrCode, msg string, printStack bool, params ...interface{}) *Error {
	return &Error{
		code:       code,
		msg:        msg,
		context:    params,
		printStack: printStack,
	}
}

func PanicSysErr(errMsg string, params ...interface{}) {
	panic(NewError(GenericErr, errMsg, true, params...))
}

func PanicSysErrNonNil(err error, errMsg string, params ...interface{}) {
	if err != nil {
		params = append(params, err)
		panic(NewError(GenericErr, errMsg, true, params...))
	}
}

func PanicBizErrNonNil(err error, errMsg string, params ...interface{}) {
	if err != nil {
		params = append(params, err)
		panic(NewError(GenericErr, errMsg, false, params...))
	}
}

func PanicBizErr(errMsg string, params ...interface{}) {
	panic(NewError(GenericErr, errMsg, false, params...))
}
