package peerwire

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRquestMetadataFiles(t *testing.T) {

	data := withExtendedFetchMetaMsg(3, ExRequest, 0)
	req := parseExtendedFetchMetaMsg(data)
	assert.Equal(t, 3, req.ExMessageId())
	assert.Equal(t, ExRequest, req.MsgType())
	assert.Equal(t, 0, req.PieceNum())
}

func TestRquestMetadataFilesWithData(t *testing.T) {

	data := withExtendedFetchMetaMsg(3, ExData, 0, []byte{0x00, 0x01})
	req := parseExtendedFetchMetaMsg(data)
	assert.Equal(t, 3, req.ExMessageId())
	assert.Equal(t, ExData, req.MsgType())
	assert.Equal(t, 0, req.PieceNum())
	assert.Equal(t, []byte{0x00, 0x01}, req.Data())
}
