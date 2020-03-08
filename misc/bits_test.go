package misc

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXor(t *testing.T) {
	assert.Equal(t, hexDecode("020100"), Xor(hexDecode("010002"), hexDecode("030102")))
	assert.Equal(t, hexDecode("FFFFFF"), Xor(hexDecode("000000"), hexDecode("FFFFFF")))
}

func hexDecode(src string) []byte {
	bytes, _ := hex.DecodeString(src)
	return bytes
}
