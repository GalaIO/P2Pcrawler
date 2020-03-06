package dht

import (
	"github.com/GalaIO/P2Pcrawler/misc"
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
	testList := misc.Dict{
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

	testList := misc.Dict{
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

	testList := misc.Dict{
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

	res, _ := EncodeSlice(misc.List{"spam", "eggs"})
	assert.Equal(t, "l4:spam4:eggse", res)

	res, _ = EncodeSlice(misc.List{10, "eggs"})
	assert.Equal(t, "li10e4:eggse", res)

	res, _ = EncodeSlice(misc.List{-10, 0})
	assert.Equal(t, "li-10ei0ee", res)

	res, _ = EncodeSlice(misc.List{})
	assert.Equal(t, "le", res)

	_, err := EncodeSlice(misc.List{1.0})
	assert.Equal(t, WrongDecodeParamErr, err)

	res, _ = EncodeSlice(misc.List{10, "abc", misc.List{"a", "b"}})
	assert.Equal(t, "li10e3:abcl1:a1:bee", res)
}

func TestListDecode(t *testing.T) {

	var res misc.List
	var err error

	res, _ = DecodeSlice("l4:spam4:eggse")
	assert.Equal(t, misc.List{"spam", "eggs"}, res)

	res, _ = DecodeSlice("li10e4:eggse")
	assert.Equal(t, misc.List{10, "eggs"}, res)

	res, _ = DecodeSlice("li-10ei0ee")
	assert.Equal(t, misc.List{-10, 0}, res)

	res, _ = DecodeSlice("li-10ei0eli1ei2eee")
	assert.Equal(t, misc.List{-10, 0, misc.List{1, 2}}, res)

	res, _ = DecodeSlice("li10ei0ei-1elei10ee")
	assert.Equal(t, misc.List{10, 0, -1, misc.List{}, 10}, res)

	res, _ = DecodeSlice("le")
	assert.Equal(t, misc.List{}, res)

	_, err = DecodeSlice("li-0e1:2222e")
	assert.Equal(t, WrongDecodeParamErr, err)

	_, err = DecodeSlice("li:e")
	assert.Equal(t, WrongDecodeParamErr, err)

	_, err = DecodeSlice("li10e")
	assert.Equal(t, WrongDecodeParamErr, err)

}

func TestMapEncode(t *testing.T) {

	res, _ := EncodeDict(misc.Dict{"cow": "moo", "spam": "eggs"})
	assert.Equal(t, "d3:cow3:moo4:spam4:eggse", res)

	res, _ = EncodeDict(misc.Dict{"name": "xiaohua", "age": 10})
	assert.Equal(t, "d4:name7:xiaohua3:agei10ee", res)

	res, _ = EncodeDict(misc.Dict{"spam": misc.List{"a", "b"}, "age": misc.List{10, 20}})
	assert.Equal(t, "d4:spaml1:a1:be3:ageli10ei20eee", res)

	res, _ = EncodeDict(misc.Dict{"spam": misc.List{"a", "b"}, "age": misc.Dict{"math": 10, "english": 20}})
	assert.Equal(t, "d4:spaml1:a1:be3:aged4:mathi10e7:englishi20eee", res)

}

func TestMapDecode(t *testing.T) {

	var res misc.Dict
	var err error

	res, _ = DecodeDict("d3:cow3:moo4:spam4:eggse")
	assert.Equal(t, misc.Dict{"cow": "moo", "spam": "eggs"}, res)

	res, _ = DecodeDict("d4:name7:xiaohua3:agei10ee")
	assert.Equal(t, misc.Dict{"name": "xiaohua", "age": 10}, res)

	_, err = DecodeDict("d1:name7:xiaohua:agei10ee")
	assert.Equal(t, WrongDecodeParamErr, err)

	res, _ = DecodeDict("d4:spaml1:a1:be3:ageli10ei20eee")
	assert.Equal(t, misc.Dict{"spam": misc.List{"a", "b"}, "age": misc.List{10, 20}}, res)

	res, _ = DecodeDict("d4:spaml1:a1:be3:aged4:mathi10e7:englishi20eee")
	assert.Equal(t, misc.Dict{"spam": misc.List{"a", "b"}, "age": misc.Dict{"math": 10, "english": 20}}, res)

}
