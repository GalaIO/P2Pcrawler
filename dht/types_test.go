package dht

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictEncode(t *testing.T) {
	d := &Dict{"name": "xiaohua", "grade": List{100, 99}, "family": Dict{"sister": "xiaomei"}}
	res, _ := d.benEncode()
	assert.Equal(t, "d4:name7:xiaohua5:gradeli100ei99ee6:familyd6:sister7:xiaomeiee", res)
}
