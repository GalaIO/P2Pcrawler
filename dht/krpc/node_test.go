package krpc

import (
	"encoding/hex"
	"github.com/GalaIO/P2Pcrawler/misc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNodeInfoParse(t *testing.T) {
	src := "663978af8cc93057aa391db6c02e28920846507505bdb939c8ea66da681d446a55f0c05b972c4464feb2c006165118157d53e120254ecff1f1bbe9ebb3a6db3c870c3e99245e0d90c000945c896521fe13a6a764cfe9057c2bf872baabf97602fcdbd5884fee82c236b0eee0f42a2abba9a3af85282372e10124858905bdbb5ac9290c0edce5a05d0b3d275b9557f3e66ca9b7b00c126d7b29c16bf70f20e449f1f1bbe9ebb3a6db3c870c3e99245e52640b1fd553d7075c3174989b29a592f287cc5095033fded997052d0ae6a71ae1"
	bytes, _ := hex.DecodeString(src)
	nodes := parseNodeInfo(string(bytes))
	for _, v := range nodes {
		t.Log("id", hex.EncodeToString([]byte(v.Id)), "host", v.Addr.String())
	}
	infos := joinNodeInfos(nodes)
	assert.Equal(t, src, hex.EncodeToString([]byte(infos)))
}

func TestAddrParse(t *testing.T) {
	vals := misc.List{"1db6c0"}
	addrs := parsePeerInfo(vals)
	assert.Equal(t, vals, joinPeerInfos(addrs))
}
