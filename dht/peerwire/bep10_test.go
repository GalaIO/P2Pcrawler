package peerwire

import (
	"github.com/GalaIO/P2Pcrawler/misc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenExtendedHandShakeNoSize(t *testing.T) {
	data := withExtendedhandShake(misc.Dict{"ut_metadata": 3}, nil)
	exHandShakeMsg := parseExtendedHandShake(data)
	meta := exHandShakeMsg.Metadata()
	assert.Equal(t, 3, meta.GetInteger("ut_metadata"))
}

func TestGenExtendedHandShake(t *testing.T) {
	data := withExtendedhandShake(misc.Dict{"ut_metadata": 3}, misc.Dict{"metadata_size": 32000})
	exHandShakeMsg := parseExtendedHandShake(data)
	meta := exHandShakeMsg.Metadata()
	dict := exHandShakeMsg.Dict()
	assert.Equal(t, 3, meta.GetInteger("ut_metadata"))
	assert.Equal(t, 32000, dict.GetInteger("metadata_size"))
}
