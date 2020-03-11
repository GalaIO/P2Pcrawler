package test

import (
	"encoding/hex"
	"encoding/json"
	"github.com/GalaIO/P2Pcrawler/dht/peerwire"
	"github.com/GalaIO/P2Pcrawler/misc"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func TestFetchMeatFile(t *testing.T) {
	laddr := "14.131.93.231:12951"
	infoHash, _ := hex.DecodeString("9bd56482d6fd6a436f5051d3b9560cdd942a5962")
	metaData, err := peerwire.FetchMetaData(laddr, peerwire.GeneratePeerId("test1"), infoHash)
	assert.Equal(t, nil, err)
	t.Log(string(metaData))
}

func TestParseAddrAndFetch(t *testing.T) {
	addrs, infos := parseAddr()
	for i, laddr := range addrs {
		infoHash, _ := hex.DecodeString(infos[i])
		metaData, err := peerwire.FetchMetaData(laddr, peerwire.GeneratePeerId("test1"), infoHash)
		assert.Equal(t, nil, err)
		t.Log(string(metaData))
	}
}

func parseAddr() ([]string, []string) {
	raw := `{"level":"info","ts":1583851453.4604506,"msg":"announce_peer","port":7023,"infoHash":"9be16a73baa0fad31731f5d57ddf5e61a9891a67","addr":"114.83.167.167:7565"}
{"level":"info","ts":1583851464.1019855,"msg":"announce_peer","addr":"223.198.111.85:2872","port":26851,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583851465.9373474,"msg":"announce_peer","addr":"125.62.6.234:2890","port":49397,"infoHash":"9b69895fda0ff29b20e7fe227bd8bf29173cd750"}
{"level":"info","ts":1583851468.6364973,"msg":"announce_peer","addr":"101.240.164.109:7885","port":7885,"infoHash":"9b78c1e659270dc10bad77b698db2e3d90755900"}
{"level":"info","ts":1583851470.2525985,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"117.151.41.96:6831","port":15000}
{"level":"info","ts":1583851470.4481056,"msg":"announce_peer","addr":"123.234.142.16:8581","port":8581,"infoHash":"9be16a73baa0fad31731f5d57ddf5e61a9891a67"}
{"level":"info","ts":1583851472.583141,"msg":"announce_peer","addr":"124.23.134.102:6755","port":41485,"infoHash":"9bcf45b1147e88a0e30ae3fe84e7b608bb384549"}
{"level":"info","ts":1583851475.1251445,"msg":"announce_peer","addr":"188.232.172.155:26982","port":26982,"infoHash":"9be168339c4269a938ebfb2da1679257b5d48a52"}
{"level":"info","ts":1583851475.579339,"msg":"announce_peer","addr":"36.48.109.7:8603","port":15000,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583851479.3891997,"msg":"announce_peer","port":15000,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"121.232.149.98:6881"}
{"level":"info","ts":1583851480.0350847,"msg":"announce_peer","addr":"114.86.217.64:6881","port":15000,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583851480.8829734,"msg":"announce_peer","addr":"180.78.99.127:20535","port":8447,"infoHash":"9b4516671985d5087b08376eb5cea5c7f58bd6f8"}
{"level":"info","ts":1583851487.4144604,"msg":"announce_peer","addr":"118.197.17.213:21043","port":9601,"infoHash":"9bffb4694bca866764a1e53f7ae567f4f723bdf4"}
{"level":"info","ts":1583851487.8447099,"msg":"announce_peer","addr":"180.78.99.127:20535","port":8447,"infoHash":"9bd4def028362ea196a803a3df47881250452ccc"}
{"level":"info","ts":1583851488.76618,"msg":"announce_peer","addr":"223.192.198.136:7875","port":9655,"infoHash":"9bc3363dec34daeefab679739bb7fb56fda3c145"}
{"level":"info","ts":1583851491.1521413,"msg":"announce_peer","addr":"1.95.3.165:26621","port":32149,"infoHash":"9bd9775ab8310133eabc7b7181e8850ad164d07b"}
{"level":"info","ts":1583851495.0991547,"msg":"announce_peer","addr":"59.33.49.145:9373","port":57146,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583851497.4976947,"msg":"announce_peer","addr":"1.95.215.45:29319","port":29319,"infoHash":"9b2d8da87cf7d4817947af217b6a50afc20055b8"}
{"level":"info","ts":1583851497.5928946,"msg":"announce_peer","addr":"124.203.128.212:24420","port":24420,"infoHash":"9bddc6df6a096a2946b68d7c8c34c5ce29d60162"}
{"level":"info","ts":1583851502.1080844,"msg":"announce_peer","addr":"111.182.106.226:23249","port":8699,"infoHash":"9be16a73baa0fad31731f5d57ddf5e61a9891a67"}
{"level":"info","ts":1583851502.1091151,"msg":"announce_peer","addr":"223.192.199.170:63161","port":8325,"infoHash":"9be16a73baa0fad31731f5d57ddf5e61a9891a67"}
{"level":"info","ts":1583851502.5201058,"msg":"announce_peer","addr":"42.197.199.26:1642","port":7977,"infoHash":"9bb6330b35e77fee9a6e4208da2f3c4b71904f93"}
{"level":"info","ts":1583851507.38631,"msg":"announce_peer","addr":"42.197.10.71:25168","port":8002,"infoHash":"9be5b82521110071f12277090988547f2f57c569"}
{"level":"info","ts":1583851507.694756,"msg":"announce_peer","port":13986,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"123.122.48.32:13986"}
{"level":"info","ts":1583851509.2470975,"msg":"announce_peer","infoHash":"9be168339c4269a938ebfb2da1679257b5d48a52","addr":"176.49.87.219:27484","port":27484}
{"level":"info","ts":1583851511.0150542,"msg":"announce_peer","addr":"42.197.199.26:1642","port":7977,"infoHash":"9bda59c0af2701e52972ee7c7f49ceba5442b60f"}
{"level":"info","ts":1583851515.1962967,"msg":"announce_peer","infoHash":"9be1786a0c3590790895caed1b9ff995a33ed391","addr":"180.114.90.109:51787","port":8317}
{"level":"info","ts":1583851516.8454366,"msg":"announce_peer","addr":"117.143.113.69:6791","port":14718,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583851519.1081402,"msg":"announce_peer","addr":"39.109.214.163:16437","port":16437,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583851530.300037,"msg":"announce_peer","addr":"117.32.92.105:8969","port":27255,"infoHash":"9be16a73baa0fad31731f5d57ddf5e61a9891a67"}
{"level":"info","ts":1583851530.9162614,"msg":"announce_peer","addr":"49.220.134.248:20522","port":9387,"infoHash":"9bf676a4ce7760c0f901a540f114093f0a99fdf6"}
{"level":"info","ts":1583851530.9447393,"msg":"announce_peer","port":65535,"infoHash":"9becfac744d7a09abf15d2e11b1307bd70ca8ae2","addr":"115.34.34.64:65535"}`
	rawList := strings.Split(raw, "\n")
	addrList := make([]string, len(rawList))
	infoHashList := make([]string, len(rawList))
	for i, v := range rawList {
		dict := misc.Dict{}
		err := json.Unmarshal([]byte(v), &dict)
		if err != nil {
			log.Fatal(err)
		}
		addrList[i] = dict.GetString("addr")
		infoHashList[i] = dict.GetString("infoHash")
		log.Println(addrList[i], infoHashList[i])
	}
	return addrList, infoHashList
}
