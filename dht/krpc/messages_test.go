package krpc

import (
	"github.com/GalaIO/P2Pcrawler/misc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenErrMsg(t *testing.T) {

	var err Response

	err = WithErr("aa", 202, "test protocol err")
	assert.Equal(t, misc.Dict{"t": "aa", "y": "e", "e": misc.List{202, "test protocol err"}}, err.RawData())

	err = WithParamErr("aa", "test param err")
	assert.Equal(t, misc.Dict{"t": "aa", "y": "e", "e": misc.List{203, "test param err"}}, err.RawData())
}
