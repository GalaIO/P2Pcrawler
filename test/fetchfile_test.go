package test

import (
	"encoding/hex"
	"github.com/GalaIO/P2Pcrawler/dht/peerwire"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchMeatFile(t *testing.T) {
	//laddr := "14.131.93.231:12951"
	laddr := "87.98.162.88:6881"
	infoHash, _ := hex.DecodeString("9be1701a5ac5a93f541b11a8e2fa24b562fa66bb")
	metaData, err := peerwire.FetchMetaData(laddr, infoHash)
	assert.Equal(t, nil, err)
	t.Log(string(metaData))
}
