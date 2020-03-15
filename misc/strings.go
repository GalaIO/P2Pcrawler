package misc

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"unsafe"
)

func ToString(v interface{}) string {
	if vs, ok := v.(string); ok {
		return vs
	}
	if vs, ok := v.(error); ok {
		return vs.Error()
	}
	if vs, ok := v.(fmt.Stringer); ok {
		return vs.String()
	}
	bytes, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return Bytes2Str(bytes)
}

func Str2Hex(str string) string {
	return hex.EncodeToString(Str2Bytes(str))
}

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
