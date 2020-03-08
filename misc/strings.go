package misc

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
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
	return string(bytes)
}

func Str2Hex(str string) string {
	return hex.EncodeToString([]byte(str))
}
