package dht

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrMsg(t *testing.T) {

	assert.Equal(t, 200, int(NoneErr))
	assert.Equal(t, 201, int(GenericErr))
	assert.Equal(t, 204, int(UnknowMethod))
}
