package peerwire

import (
	"encoding/hex"
	"github.com/GalaIO/P2Pcrawler/misc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenExtendedHandShakeNoSize(t *testing.T) {
	data := withExtendedhandShake(misc.Dict{"ut_metadata": 1}, nil)
	assert.Equal(t, "0000001a140064313a6d6431313a75745f6d657461646174616931656565", hex.EncodeToString(data))
	exHandShakeMsg := parseExtendedHandShake(data[4:])
	meta := exHandShakeMsg.Metadata()
	assert.Equal(t, 1, meta.GetInteger("ut_metadata"))
}

func TestParseExtendedHandShake(t *testing.T) {
	bytes, _ := hex.DecodeString("140064343a69707634343ac0a8016a313a6d6431313a75745f6d65746164617461693165363a75745f7065786932656531333a6d657461646174615f73697a6569363938373765313a7069313239353165343a726571716935313265313a7631363a4269745370697269742076332e362e3065")
	exHandShakeMsg := parseExtendedHandShake(bytes)
	meta := exHandShakeMsg.Metadata()
	dict := exHandShakeMsg.Dict()
	assert.Equal(t, 1, meta.GetInteger("ut_metadata"))
	assert.Equal(t, 2, meta.GetInteger("ut_pex"))
	assert.Equal(t, 69877, dict.GetInteger("metadata_size"))
	assert.Equal(t, 12951, dict.GetInteger("p"))
	assert.Equal(t, 512, dict.GetInteger("reqq"))
	assert.Equal(t, "BitSpirit v3.6.0", dict.GetString("v"))
}

func TestGenExtendedHandShake(t *testing.T) {
	data := withExtendedhandShake(misc.Dict{"ut_metadata": 3}, misc.Dict{"metadata_size": 32000})
	exHandShakeMsg := parseExtendedHandShake(data[4:])
	meta := exHandShakeMsg.Metadata()
	dict := exHandShakeMsg.Dict()
	assert.Equal(t, 3, meta.GetInteger("ut_metadata"))
	assert.Equal(t, 32000, dict.GetInteger("metadata_size"))
}
