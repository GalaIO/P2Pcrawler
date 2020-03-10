package misc

import (
	"encoding/hex"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestIntergeEncode(t *testing.T) {
	testList := map[int]interface{}{
		3:  "i3e",
		0:  "i0e",
		-3: "i-3e",
	}
	for src := range testList {
		res, _ := EncodeInteger(src)
		assert.Equal(t, testList[src], res)
	}
}
func TestIntergeDecode(t *testing.T) {
	testList := Dict{
		"i3e":   3,
		"i0e":   0,
		"i-3e":  -3,
		"":      WrongDecodeParamErr,
		"-3e":   WrongDecodeParamErr,
		"i03e":  WrongDecodeParamErr,
		"i00e":  WrongDecodeParamErr,
		"i1.0e": WrongDecodeParamErr,
		"ife":   WrongDecodeParamErr,
		"i-0e":  WrongDecodeParamErr,
	}
	for src := range testList {
		res, err := DecodeInteger(src)
		if err != nil {
			assert.Equal(t, testList[src], err)
			continue
		}
		assert.Equal(t, testList[src], res)
	}
}

func TestStringEncode(t *testing.T) {

	testList := Dict{
		"spam":  "4:spam",
		"hello": "5:hello",
		"":      "0:",
	}
	for src := range testList {
		res, _ := EncodeString(src)
		assert.Equal(t, testList[src], res)
	}
}

func TestStringDecode(t *testing.T) {

	testList := Dict{
		"4:spam":  "spam",
		"5:hello": "hello",
		"0:":      "",
		"10:":     WrongDecodeParamErr,
		"0:spam":  WrongDecodeParamErr,
		"0spam":   WrongDecodeParamErr,
		":spam":   WrongDecodeParamErr,
	}
	for src := range testList {
		res, err := DecodeString(src)
		if err != nil {
			assert.Equal(t, testList[src], err)
			continue
		}
		assert.Equal(t, testList[src], res)
	}
}

func TestListEncode(t *testing.T) {

	res, _ := EncodeSlice(List{"spam", "eggs"})
	assert.Equal(t, "l4:spam4:eggse", res)

	res, _ = EncodeSlice(List{10, "eggs"})
	assert.Equal(t, "li10e4:eggse", res)

	res, _ = EncodeSlice(List{-10, 0})
	assert.Equal(t, "li-10ei0ee", res)

	res, _ = EncodeSlice(List{})
	assert.Equal(t, "le", res)

	_, err := EncodeSlice(List{1.0})
	assert.Equal(t, WrongDecodeParamErr, err)

	res, _ = EncodeSlice(List{10, "abc", List{"a", "b"}})
	assert.Equal(t, "li10e3:abcl1:a1:bee", res)
}

func TestListDecode(t *testing.T) {

	var res List
	var err error

	res, _ = DecodeSlice("l4:spam4:eggse")
	assert.Equal(t, List{"spam", "eggs"}, res)

	res, _ = DecodeSlice("li10e4:eggse")
	assert.Equal(t, List{10, "eggs"}, res)

	res, _ = DecodeSlice("li-10ei0ee")
	assert.Equal(t, List{-10, 0}, res)

	res, _ = DecodeSlice("li-10ei0eli1ei2eee")
	assert.Equal(t, List{-10, 0, List{1, 2}}, res)

	res, _ = DecodeSlice("li10ei0ei-1elei10ee")
	assert.Equal(t, List{10, 0, -1, List{}, 10}, res)

	res, _ = DecodeSlice("le")
	assert.Equal(t, List{}, res)

	_, err = DecodeSlice("li-0e1:2222e")
	assert.Equal(t, WrongDecodeParamErr, err)

	_, err = DecodeSlice("li:e")
	assert.Equal(t, WrongDecodeParamErr, err)

	_, err = DecodeSlice("li10e")
	assert.Equal(t, WrongDecodeParamErr, err)

}

func TestMapEncode(t *testing.T) {

	res, _ := EncodeDict(Dict{"cow": "moo", "spam": "eggs"})
	assert.Equal(t, "d3:cow3:moo4:spam4:eggse", res)

	res, _ = EncodeDict(Dict{"name": "xiaohua", "age": 10})
	assert.Equal(t, "d4:name7:xiaohua3:agei10ee", res)

	res, _ = EncodeDict(Dict{"spam": List{"a", "b"}, "age": List{10, 20}})
	assert.Equal(t, "d4:spaml1:a1:be3:ageli10ei20eee", res)

	res, _ = EncodeDict(Dict{"spam": List{"a", "b"}, "age": Dict{"math": 10, "english": 20}})
	assert.Equal(t, "d4:spaml1:a1:be3:aged4:mathi10e7:englishi20eee", res)

}

func TestMapDecode(t *testing.T) {

	var res Dict
	var err error

	res, _ = DecodeDict("d3:cow3:moo4:spam4:eggse")
	assert.Equal(t, Dict{"cow": "moo", "spam": "eggs"}, res)

	res, _ = DecodeDict("d4:name7:xiaohua3:agei10ee")
	assert.Equal(t, Dict{"name": "xiaohua", "age": 10}, res)

	_, err = DecodeDict("d1:name7:xiaohua:agei10ee")
	assert.Equal(t, WrongDecodeParamErr, err)

	res, _ = DecodeDict("d4:spaml1:a1:be3:ageli10ei20eee")
	assert.Equal(t, Dict{"spam": List{"a", "b"}, "age": List{10, 20}}, res)

	res, _ = DecodeDict("d4:spaml1:a1:be3:aged4:mathi10e7:englishi20eee")
	assert.Equal(t, Dict{"spam": List{"a", "b"}, "age": Dict{"math": 10, "english": 20}}, res)

}

func TestName(t *testing.T) {
	s, _ := hex.DecodeString("64313a71393a66696e645f6e6f6465313a7264323a696432303a9be1701a5ac5a93f2c945f9bbb33d224742ac3fa323a6970343a7a33a950373a6e6f646573007632363a9be1696018a69ba161651978c76c46f34a6aba747a33a950520865313a74343a30346430313a79313a7265")
	dicts, _ := DecodeDict(string(s))
	t.Log(dicts)
}
