package dht

import "testing"
import "github.com/stretchr/testify/assert"

func TestIntergeEncode(t *testing.T) {
	testList := map[int]interface{}{
		3:  "i3e",
		0:  "i0e",
		-3: "i-3e",
	}
	for src := range testList {
		res, _ := encodeInteger(src)
		assert.Equal(t, testList[src], res)
	}
}
func TestIntergeDecode(t *testing.T) {
	testList := map[string]interface{}{
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
		res, err := decodeInteger(src)
		if err != nil {
			assert.Equal(t, testList[src], err)
			continue
		}
		assert.Equal(t, testList[src], res)
	}
}

func TestStringEncode(t *testing.T) {

	testList := map[string]interface{}{
		"spam":  "4:spam",
		"hello": "5:hello",
		"":      "0:",
	}
	for src := range testList {
		res, _ := encodeString(src)
		assert.Equal(t, testList[src], res)
	}
}

func TestStringDecode(t *testing.T) {

	testList := map[string]interface{}{
		"4:spam":  "spam",
		"5:hello": "hello",
		"0:":      "",
		"10:":     WrongDecodeParamErr,
		"0:spam":  WrongDecodeParamErr,
		"0spam":   WrongDecodeParamErr,
		":spam":   WrongDecodeParamErr,
	}
	for src := range testList {
		res, err := decodeString(src)
		if err != nil {
			assert.Equal(t, testList[src], err)
			continue
		}
		assert.Equal(t, testList[src], res)
	}
}

func TestListEncode(t *testing.T) {

	res, _ := encodeSlice([]interface{}{"spam", "eggs"})
	assert.Equal(t, "l4:spam4:eggse", res)

	res, _ = encodeSlice([]interface{}{10, "eggs"})
	assert.Equal(t, "li10e4:eggse", res)

	res, _ = encodeSlice([]interface{}{-10, 0})
	assert.Equal(t, "li-10ei0ee", res)

	res, _ = encodeSlice([]interface{}{})
	assert.Equal(t, "le", res)

	_, err := encodeSlice([]interface{}{1.0})
	assert.Equal(t, WrongDecodeParamErr, err)

	res, _ = encodeSlice([]interface{}{10, "abc", []interface{}{"a", "b"}})
	assert.Equal(t, "li10e3:abcl1:a1:bee", res)
}

func TestListDecode(t *testing.T) {

	var res []interface{}
	var err error

	res, _ = decodeSlice("l4:spam4:eggse")
	assert.Equal(t, []interface{}{"spam", "eggs"}, res)

	res, _ = decodeSlice("li10e4:eggse")
	assert.Equal(t, []interface{}{10, "eggs"}, res)

	res, _ = decodeSlice("li-10ei0ee")
	assert.Equal(t, []interface{}{-10, 0}, res)

	res, _ = decodeSlice("li-10ei0eli1ei2eee")
	assert.Equal(t, []interface{}{-10, 0, []interface{}{1, 2}}, res)

	res, _ = decodeSlice("li10ei0ei-1elei10ee")
	assert.Equal(t, []interface{}{10, 0, -1, []interface{}{}, 10}, res)

	res, _ = decodeSlice("le")
	assert.Equal(t, []interface{}{}, res)

	_, err = decodeSlice("li-0e1:2222e")
	assert.Equal(t, WrongDecodeParamErr, err)

	_, err = decodeSlice("li:e")
	assert.Equal(t, WrongDecodeParamErr, err)

	_, err = decodeSlice("li10e")
	assert.Equal(t, WrongDecodeParamErr, err)

}

func TestMapEncode(t *testing.T) {

	res, _ := encodeDict(map[string]interface{}{"cow": "moo", "spam": "eggs"})
	assert.Equal(t, "d3:cow3:moo4:spam4:eggse", res)

	res, _ = encodeDict(map[string]interface{}{"name": "xiaohua", "age": 10})
	assert.Equal(t, "d4:name7:xiaohua3:agei10ee", res)

	res, _ = encodeDict(map[string]interface{}{"spam": []interface{}{"a", "b"}, "age": []interface{}{10, 20}})
	assert.Equal(t, "d4:spaml1:a1:be3:ageli10ei20eee", res)

	res, _ = encodeDict(map[string]interface{}{"spam": []interface{}{"a", "b"}, "age": map[string]interface{}{"math": 10, "english": 20}})
	assert.Equal(t, "d4:spaml1:a1:be3:aged4:mathi10e7:englishi20eee", res)

}

func TestMapDecode(t *testing.T) {

	var res map[string]interface{}
	var err error

	res, _ = decodeDict("d3:cow3:moo4:spam4:eggse")
	assert.Equal(t, map[string]interface{}{"cow": "moo", "spam": "eggs"}, res)

	res, _ = decodeDict("d4:name7:xiaohua3:agei10ee")
	assert.Equal(t, map[string]interface{}{"name": "xiaohua", "age": 10}, res)

	_, err = decodeDict("d1:name7:xiaohua:agei10ee")
	assert.Equal(t, WrongDecodeParamErr, err)

	res, _ = decodeDict("d4:spaml1:a1:be3:ageli10ei20eee")
	assert.Equal(t, map[string]interface{}{"spam": []interface{}{"a", "b"}, "age": []interface{}{10, 20}}, res)

	res, _ = decodeDict("d4:spaml1:a1:be3:aged4:mathi10e7:englishi20eee")
	assert.Equal(t, map[string]interface{}{"spam": []interface{}{"a", "b"}, "age": map[string]interface{}{"math": 10, "english": 20}}, res)

}
