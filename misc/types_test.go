package misc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictGetter(t *testing.T) {

	d := &Dict{"name": "xiaohua", "total": 199, "grade": List{100, 99}, "family": Dict{"sister": "xiaomei"}}

	assert.Equal(t, 199, d.GetInteger("total"))
	assert.Equal(t, "xiaohua", d.GetString("name"))
	assert.Equal(t, List{100, 99}, d.GetList("grade"))
	assert.Equal(t, Dict{"sister": "xiaomei"}, d.GetDict("family"))
}

func TestDictGetterPanic(t *testing.T) {

	d := &Dict{"name": "xiaohua", "total": 199, "grade": List{100, 99}, "family": Dict{"sister": "xiaomei"}}

	testDataMap := []struct {
		input  string
		tp     string
		output interface{}
	}{
		{
			"totalGrade",
			"int",
			"201:cannot find totalGrade's val",
		},
		{
			"name",
			"int",
			"201:name's val cannot match type of int",
		},
		{
			"total",
			"string",
			"201:total's val cannot match type of string",
		},
		{
			"name",
			"list",
			"201:name's val cannot match type of List",
		},
		{
			"name",
			"Dict",
			"201:name's val cannot match type of Dict",
		},
	}

	for _, td := range testDataMap {
		func() {
			defer func() {
				if err := recover(); err != nil {
					dhtError := err.(*Error)
					assert.Equal(t, td.output, dhtError.Error())
				}
			}()

			switch td.tp {
			case "int":
				d.GetInteger(td.input)
			case "string":
				d.GetString(td.input)
			case "List":
				d.GetList(td.input)
			case "Dict":
				d.GetDict(td.input)
			}
		}()
	}
}

func TestListGetter(t *testing.T) {

	l := &List{10, List{"spam", "eggs"}, "age", Dict{"name": "xiaoguo"}}

	assert.Equal(t, 10, l.GetInteger(0))
	assert.Equal(t, "age", l.GetString(2))
	assert.Equal(t, List{"spam", "eggs"}, l.GetList(1))
	assert.Equal(t, Dict{"name": "xiaoguo"}, l.GetDict(3))
}

func TestListGetterPanic(t *testing.T) {

	l := &List{10, List{"spam", "eggs"}, "age", Dict{"name": "xiaoguo"}}

	testDataMap := []struct {
		input  int
		tp     string
		output interface{}
	}{
		{
			1,
			"int",
			"201:val cannot match type of int",
		},
		{
			0,
			"string",
			"201:val cannot match type of string",
		},
		{
			0,
			"list",
			"201:val cannot match type of List",
		},
		{
			0,
			"Dict",
			"201:val cannot match type of Dict",
		},
	}

	for _, td := range testDataMap {
		func() {
			defer func() {
				if err := recover(); err != nil {
					dhtError := err.(*Error)
					assert.Equal(t, td.output, dhtError.Error())
				}
			}()

			switch td.tp {
			case "int":
				l.GetInteger(td.input)
			case "string":
				l.GetString(td.input)
			case "List":
				l.GetList(td.input)
			case "Dict":
				l.GetDict(td.input)
			}
		}()
	}
}
