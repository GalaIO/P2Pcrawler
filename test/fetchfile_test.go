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
	raw := `{"level":"info","ts":1583812311.1081028,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812312.3242745,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812312.3902695,"msg":"announce_peer","addr":"113.103.208.81:18067","port":16662,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812313.3121521,"msg":"announce_peer","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"125.231.89.159:18568"}
{"level":"info","ts":1583812314.388103,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812314.8159583,"msg":"announce_peer","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"125.231.89.159:18568"}
{"level":"info","ts":1583812315.268358,"msg":"announce_peer","addr":"36.228.139.229:27554","port":27554,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812315.273016,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812316.0904758,"msg":"announce_peer","addr":"125.111.47.186:1057","port":45984,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812316.3364108,"msg":"announce_peer","port":27612,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"119.246.19.137:27612"}
{"level":"info","ts":1583812316.7108464,"msg":"announce_peer","port":27554,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"36.228.139.229:27554"}
{"level":"info","ts":1583812316.7680182,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812317.1200073,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812317.4600902,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812317.9072635,"msg":"announce_peer","addr":"220.132.95.218:20804","port":20804,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812319.08052,"msg":"announce_peer","infoHash":"9a3c43dfa4194659a068b06347a82de81f7efe29","addr":"182.51.77.189:10802","port":27017}
{"level":"info","ts":1583812319.172062,"msg":"announce_peer","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"125.231.89.159:18568"}
{"level":"info","ts":1583812319.8737533,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"184.64.116.183:17270","port":17270}
{"level":"info","ts":1583812320.1992986,"msg":"announce_peer","addr":"113.88.80.159:1036","port":53898,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812320.2893999,"msg":"announce_peer","addr":"36.228.139.229:27554","port":27554,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812320.7640429,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812321.2658675,"msg":"announce_peer","addr":"36.228.139.229:27554","port":27554,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812321.270894,"msg":"announce_peer","addr":"121.69.7.42:13042","port":13042,"infoHash":"9e594ff4283a689f47d095d5ab96484e9546ddd2"}
{"level":"info","ts":1583812322.2844296,"msg":"announce_peer","addr":"36.228.139.229:27554","port":27554,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812322.9232008,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812323.9115665,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"1.170.4.192:26990","port":26990}
{"level":"info","ts":1583812324.453623,"msg":"announce_peer","addr":"223.192.199.170:4652","port":8325,"infoHash":"9be16a73baa0fad31731f5d57ddf5e61a9891a67"}
{"level":"info","ts":1583812324.9080276,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812325.0475192,"msg":"announce_peer","addr":"125.62.6.109:2161","port":20275,"infoHash":"9aba304ab2286c6719aa672688b09b91741b68ee"}
{"level":"info","ts":1583812325.526899,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812325.5596318,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812326.305193,"msg":"announce_peer","port":23467,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"184.177.64.114:23467"}
{"level":"info","ts":1583812326.643189,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812326.6438162,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812326.7630572,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812326.7787864,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812326.9168432,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812326.9502497,"msg":"announce_peer","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"124.77.232.78:23971"}
{"level":"info","ts":1583812327.0599701,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812327.1610942,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812327.2783575,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812327.4346576,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812328.436036,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812329.0639985,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812329.3407695,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812329.391447,"msg":"announce_peer","addr":"119.246.19.137:27612","port":27612,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812329.4555004,"msg":"announce_peer","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"1.170.4.192:26990"}
{"level":"info","ts":1583812329.466657,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812330.1226225,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812330.1829417,"msg":"announce_peer","addr":"125.206.82.229:31708","port":31708,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812330.2405083,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812330.33927,"msg":"announce_peer","infoHash":"9be94f66cee56c3d63a32aa9554bb9ff4150cb1a","addr":"183.227.121.126:24576","port":45062}
{"level":"info","ts":1583812330.575985,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812330.6337547,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"1.170.4.192:26990","port":26990}
{"level":"info","ts":1583812330.7473383,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812330.7832348,"msg":"announce_peer","addr":"197.185.104.200:33184","port":37649,"infoHash":"9be1751a9caf2e0e26e241f43688bd4779583993"}
{"level":"info","ts":1583812331.0768073,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812331.590784,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"1.170.4.192:26990","port":26990}
{"level":"info","ts":1583812331.8220537,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812332.0249052,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812332.4142044,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812332.5405588,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"124.77.232.78:23971","port":23971}
{"level":"info","ts":1583812333.467484,"msg":"announce_peer","infoHash":"9b830709afe4bd8f0686ee62a70a42f9f57806ba","addr":"118.29.251.110:9459","port":9459}
{"level":"info","ts":1583812333.6911206,"msg":"announce_peer","addr":"36.228.139.229:27554","port":27554,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812333.8021438,"msg":"announce_peer","addr":"115.33.90.88:7997","port":7997,"infoHash":"9e3f3af59958e91702f6bb1cac9fbe5e632f0b43"}
{"level":"info","ts":1583812333.9282696,"msg":"announce_peer","addr":"117.81.214.189:8999","port":8999,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812333.9453714,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812334.086293,"msg":"announce_peer","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"124.77.232.78:23971"}
{"level":"info","ts":1583812334.8052442,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812334.9800751,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812335.070559,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812335.4079282,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812335.9780269,"msg":"announce_peer","addr":"36.228.139.229:27554","port":27554,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812336.061471,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812336.2840576,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812336.4780946,"msg":"announce_peer","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"124.77.232.78:23971"}
{"level":"info","ts":1583812336.5198956,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812337.0120707,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812337.9300864,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812338.5413463,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812338.6160305,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812338.624021,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812338.7257767,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"119.246.19.137:27612","port":27612}
{"level":"info","ts":1583812339.4622755,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"36.228.139.229:27554","port":27554}
{"level":"info","ts":1583812339.5360901,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812339.5496078,"msg":"announce_peer","infoHash":"9a7f7d63fdb6ba14337336f951da07db1f7c68e5","addr":"180.78.231.253:20713","port":9743}
{"level":"info","ts":1583812339.5517085,"msg":"announce_peer","port":22261,"infoHash":"9be59d184594795e62a2e09d164e7d796bbbcc2f","addr":"103.63.154.154:7368"}
{"level":"info","ts":1583812340.242681,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812340.435991,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"125.231.89.159:18568","port":18568}
{"level":"info","ts":1583812340.5707285,"msg":"announce_peer","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"124.77.232.78:23971"}
{"level":"info","ts":1583812340.9162846,"msg":"announce_peer","addr":"36.228.139.229:27554","port":27554,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812341.1661663,"msg":"announce_peer","addr":"110.243.102.121:20504","port":19131,"infoHash":"9b9cdd2ac37efe79e5cca087dc5aa94b8debc2fe"}
{"level":"info","ts":1583812341.224146,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812341.3369226,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812341.7426283,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812341.7745085,"msg":"announce_peer","addr":"42.2.0.160:17185","port":17185,"infoHash":"9be1d9be2b045bf519dde838f6e3adc42050de15"}
{"level":"info","ts":1583812341.8857791,"msg":"announce_peer","port":27554,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"36.228.139.229:27554"}
{"level":"info","ts":1583812341.988857,"msg":"announce_peer","addr":"112.252.214.168:4154","port":6881,"infoHash":"9be149e1030f789b3f090c5cf986892cc8bc3de1"}
{"level":"info","ts":1583812342.4772642,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812342.4955537,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812342.5082078,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"125.231.89.159:18568","port":18568}
{"level":"info","ts":1583812342.7920396,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"125.231.89.159:18568","port":18568}
{"level":"info","ts":1583812342.9039943,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812343.4688144,"msg":"announce_peer","addr":"119.246.19.137:27612","port":27612,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812344.0570579,"msg":"announce_peer","port":60139,"infoHash":"9bc0aceee6aa034a0cec1ccffc47ab60f244f952","addr":"113.45.212.52:60139"}
{"level":"info","ts":1583812344.1265793,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812344.1881306,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812344.8010342,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812344.8121247,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"125.231.89.159:18568","port":18568}
{"level":"info","ts":1583812345.3599954,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"125.231.89.159:18568","port":18568}
{"level":"info","ts":1583812345.852609,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812345.9758024,"msg":"announce_peer","addr":"101.243.25.93:20725","port":8993,"infoHash":"9a926a34069c9414f7f139187042e1a4564791a1"}
{"level":"info","ts":1583812345.9854949,"msg":"announce_peer","port":8993,"infoHash":"9ad15ef733c52ce461fb502404ac22e310087300","addr":"101.243.25.93:20725"}
{"level":"info","ts":1583812346.1495597,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812346.8137133,"msg":"announce_peer","addr":"114.40.45.202:19947","port":19947,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812346.8838341,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812347.6000645,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812347.7591076,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812347.9959898,"msg":"announce_peer","port":8993,"infoHash":"9bc388488270ca66737adabdc4f58ade01bdd37d","addr":"101.243.25.93:20725"}
{"level":"info","ts":1583812348.0007644,"msg":"announce_peer","addr":"101.243.25.93:20725","port":8993,"infoHash":"9ad834f4da729521fe0fa9b432a76a9219a9862a"}
{"level":"info","ts":1583812348.3377042,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"124.77.232.78:23971","port":23971}
{"level":"info","ts":1583812348.3945324,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812348.446047,"msg":"announce_peer","addr":"103.27.26.192:64801","port":51510,"infoHash":"9be458b745712b97788a7101fcda84a27fc21d3b"}
{"level":"info","ts":1583812348.7839808,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"125.231.89.159:18568","port":18568}
{"level":"info","ts":1583812348.9384675,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812348.9802477,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812349.712495,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812350.1966064,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812350.212321,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812350.4043083,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812350.5254188,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812350.7919548,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"125.231.89.159:18568","port":18568}
{"level":"info","ts":1583812350.8441918,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812350.854946,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812350.9468036,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812350.9919968,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812351.1880827,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"125.231.89.159:18568","port":18568}
{"level":"info","ts":1583812351.2446585,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812351.508067,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812351.5251477,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"124.77.232.78:23971","port":23971}
{"level":"info","ts":1583812351.7360504,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812351.9470654,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812352.0876594,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812352.4079814,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812352.4161294,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812352.4261959,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812352.7751188,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812352.7911742,"msg":"announce_peer","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"124.77.232.78:23971"}
{"level":"info","ts":1583812352.9000773,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812353.3920562,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812353.4195707,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812353.8439274,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812354.0095718,"msg":"announce_peer","addr":"124.77.232.78:23971","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812354.040238,"msg":"announce_peer","addr":"125.231.89.159:18568","port":18568,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812354.1294065,"msg":"announce_peer","addr":"1.170.4.192:26990","port":26990,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583812354.1813035,"msg":"announce_peer","port":23971,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"124.77.232.78:23971"}`
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
