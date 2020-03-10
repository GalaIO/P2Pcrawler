package peerwire

import (
	"encoding/hex"
	"github.com/GalaIO/P2Pcrawler/misc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRquestMetadataFiles(t *testing.T) {

	data := withExtendedFetchMetaMsg(1, ExRequest, 0, 0)
	assert.Equal(t, "0000001b140164383a6d73675f74797065693065353a706965636569306565", hex.EncodeToString(data))
	req := parseExtendedFetchMetaMsg(data[4:])
	assert.Equal(t, 1, req.ExMessageId())
	assert.Equal(t, ExRequest, req.MsgType())
	assert.Equal(t, 0, req.PieceNum())
}

func TestParseMetadataFiles(t *testing.T) {

	data, _ := hex.DecodeString("140164383a6d73675f74797065693165353a706965636569306531303a746f74616c5f73697a65693639383737656564333a636f77333a6d6f6f343a7370616d343a6567677365")
	req := parseExtendedFetchMetaMsg(data)
	//assert.Equal(t, 3, req.ExMessageId())
	assert.Equal(t, ExData, req.MsgType())
	assert.Equal(t, 0, req.PieceNum())
	assert.Equal(t, 69877, req.TotalSize())
	assert.True(t, SizeOf16KB >= len(req.Data()))
	dict, err := misc.DecodeDict(misc.Bytes2Str(req.Data()))
	assert.Equal(t, nil, err)
	assert.Equal(t, misc.Dict{"cow": "moo", "spam": "eggs"}, dict)
}

func TestRquestMetadataFilesWithData(t *testing.T) {

	data := withExtendedFetchMetaMsg(3, ExData, 0, 2, []byte{0x00, 0x01})
	req := parseExtendedFetchMetaMsg(data[4:])
	assert.Equal(t, 3, req.ExMessageId())
	assert.Equal(t, ExData, req.MsgType())
	assert.Equal(t, 0, req.PieceNum())
	assert.Equal(t, []byte{0x00, 0x01}, req.Data())
}
