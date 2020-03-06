package dht

import (
	"errors"
	"github.com/GalaIO/P2Pcrawler/misc"
	"reflect"
	"strconv"
	"strings"
)

var WrongDecodeParamErr = errors.New("wrong benDecode param")

func EncodeInteger(src int) (string, error) {
	return "i" + strconv.Itoa(src) + "e", nil
}

func DecodeInteger(src string) (int, error) {

	if len(src) < 2 {
		return 0, WrongDecodeParamErr
	}

	if src[0] != 'i' || src[len(src)-1] != 'e' {
		return 0, WrongDecodeParamErr
	}

	result, next, err := innerDecodeInteger(src, 0)

	if next != len(src) || err != nil {
		return 0, WrongDecodeParamErr
	}
	return result, nil
}

func innerDecodeInteger(src string, start int) (int, int, error) {

	if src[start] != 'i' {
		return 0, -1, WrongDecodeParamErr
	}

	i := start + 1
	for ; i < len(src) && src[i] != 'e'; i++ {
		// pass -0 -02
		if src[i-1] == '-' && src[i] == '0' {
			return 0, -1, WrongDecodeParamErr
		}

		// pass 00 02
		if src[i-1] == 'i' && src[i] == '0' && src[i+1] != 'e' {
			return 0, -1, WrongDecodeParamErr
		}
	}

	if i > start && src[i] == 'e' {
		result, err := strconv.Atoi(src[start+1 : i])
		if err != nil {
			return 0, -1, WrongDecodeParamErr
		}
		return result, i + 1, nil
	}

	return 0, -1, WrongDecodeParamErr
}

func EncodeString(src string) (string, error) {
	return strconv.Itoa(len(src)) + ":" + src, nil
}

func DecodeString(src string) (string, error) {

	if len(src) < 2 {
		return "", WrongDecodeParamErr
	}
	subs := strings.Split(src, ":")
	if len(subs) < 1 || len(subs) > 2 {
		return "", WrongDecodeParamErr
	}

	slen, err := strconv.Atoi(subs[0])
	if err != nil {
		return "", WrongDecodeParamErr
	}

	if len(subs[1]) != slen {
		return "", WrongDecodeParamErr
	}

	return subs[1], nil
}

func innerDecodeString(src string, start int) (string, int, error) {
	idx := indexFirstByteInStr(src, start, len(src), ':')
	if idx <= start {
		return "", -1, WrongDecodeParamErr
	}

	slen, err := strconv.Atoi(src[start:idx])
	if err != nil {
		return "", -1, err
	}

	start = idx + slen + 1

	if start >= len(src) {
		return "", -1, WrongDecodeParamErr
	}
	return src[idx+1 : start], start, nil
}

func EncodeSlice(src misc.List) (string, error) {

	if src == nil {
		return "", WrongDecodeParamErr
	}

	res := ""
	for _, item := range src {
		tmp, err := encodeItem(item)
		if err != nil {
			return "", err
		}
		res += tmp
	}

	return "l" + res + "e", nil
}

func encodeItem(item interface{}) (string, error) {
	t := reflect.TypeOf(item)
	var err error
	var tmp string
	switch t.Kind() {
	case reflect.Int:
		tmp, err = EncodeInteger(item.(int))
	case reflect.String:
		tmp, err = EncodeString(item.(string))
	case reflect.Slice:
		if ls, ok := item.(misc.List); ok {
			tmp, err = EncodeSlice(ls)
		} else {
			err = WrongDecodeParamErr
		}
	case reflect.Map:
		if ls, ok := item.(misc.Dict); ok {
			tmp, err = EncodeDict(ls)
		} else {
			err = WrongDecodeParamErr
		}
	default:
		err = WrongDecodeParamErr
	}
	return tmp, err
}

func DecodeSlice(src string) (misc.List, error) {
	if len(src) < 2 {
		return nil, WrongDecodeParamErr
	}

	if src[0] != 'l' || src[len(src)-1] != 'e' {
		return nil, WrongDecodeParamErr
	}

	result, next, err := innerDecodeSlice(src, 0)

	if err != nil {
		return nil, err
	}

	if next != len(src) {
		return nil, WrongDecodeParamErr
	}
	return result, nil
}

func innerDecodeSlice(src string, start int) (misc.List, int, error) {
	if src[start] != 'l' {
		return nil, -1, WrongDecodeParamErr
	}

	result := make(misc.List, 0, 16)
	i := start + 1
	for i < len(src) && src[i] != 'e' {
		tmp, nextIdx, err := decodeItem(src, i)
		if err != nil {
			return nil, -1, err
		}
		result = append(result, tmp)
		i = nextIdx
	}
	return result, i + 1, nil
}

func decodeItem(src string, i int) (interface{}, int, error) {
	if src[i] == 'i' {
		return innerDecodeInteger(src, i)
	} else if src[i] == 'l' {
		return innerDecodeSlice(src, i)
	} else if src[i] == 'd' {
		return innerDecodeDict(src, i)
	} else if src[i] >= '0' && src[i] <= '9' {
		return innerDecodeString(src, i)
	} else {
		return nil, -1, WrongDecodeParamErr
	}
}

func EncodeDict(src misc.Dict) (string, error) {

	if src == nil {
		return "", WrongDecodeParamErr
	}
	str := ""
	for k, v := range src {
		ktmp, err := EncodeString(k)
		if err != nil {
			return "", err
		}
		str += ktmp

		vtmp, err := encodeItem(v)
		if err != nil {
			return "", err
		}
		str += vtmp
	}
	return "d" + str + "e", nil
}

func DecodeDict(src string) (misc.Dict, error) {

	tLen := len(src)
	if tLen < 2 {
		return nil, WrongDecodeParamErr
	}

	if src[0] != 'd' || src[tLen-1] != 'e' {
		return nil, WrongDecodeParamErr
	}

	result, next, err := innerDecodeDict(src, 0)

	if next < len(src) || err != nil {
		return nil, err
	}

	return result, nil
}

func innerDecodeDict(src string, start int) (misc.Dict, int, error) {

	if src[start] != 'd' {
		return nil, -1, WrongDecodeParamErr
	}

	result := make(misc.Dict)

	i := start + 1
	for i < len(src) && src[i] != 'e' {
		ktmp, next, err := decodeItem(src, i)
		if err != nil {
			return nil, -1, err
		}

		key, ok := ktmp.(string)
		if !ok {
			return nil, -1, WrongDecodeParamErr
		}

		val, next, err := decodeItem(src, next)
		if err != nil {
			return nil, -1, err
		}

		result[key] = val
		i = next
	}

	return result, i + 1, nil
}

func indexFirstByteInStr(str string, start, end int, ch byte) int {
	for i := start; i < len(str) && i >= 0 && i <= end; i++ {
		if str[i] == ch {
			return i
		}
	}
	return -1
}
