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
	raw := `{"level":"info","ts":1583740760.381232,"msg":"announce_peer","addr":"114.26.40.207:64525","port":8740,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740760.691286,"msg":"announce_peer","addr":"114.26.40.207:64525","port":8740,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740761.184609,"msg":"announce_peer","addr":"121.177.133.63:40282","port":40282,"infoHash":"9be16925eaa1518a8fd02a9ba32935312468507a"}
{"level":"info","ts":1583740761.7409508,"msg":"announce_peer","port":8740,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"114.26.40.207:64525"}
{"level":"info","ts":1583740761.7511702,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"114.26.40.207:64525","port":8740}
{"level":"info","ts":1583740761.7650995,"msg":"announce_peer","port":13660,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"36.4.196.66:7887"}
{"level":"info","ts":1583740762.2582445,"msg":"announce_peer","addr":"112.118.183.29:1187","port":6900,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740762.289413,"msg":"announce_peer","infoHash":"9be1786a0c3590790895caed1b9ff995a33ed391","addr":"111.14.10.191:6391","port":7180}
{"level":"info","ts":1583740762.3014402,"msg":"announce_peer","port":8740,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"114.26.40.207:64525"}
{"level":"info","ts":1583740762.3804328,"msg":"announce_peer","addr":"222.248.16.131:24213","port":24213,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740762.7211435,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"114.26.40.207:64525","port":8740}
{"level":"info","ts":1583740763.0212014,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"114.26.40.207:64525","port":8740}
{"level":"info","ts":1583740763.4112682,"msg":"announce_peer","addr":"114.26.40.207:64525","port":8740,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740764.615519,"msg":"announce_peer","addr":"113.47.4.226:48992","port":48901,"infoHash":"9bff490e4c46a0e6f81d782f6cf93ba82ece2a58"}
{"level":"info","ts":1583740764.932137,"msg":"announce_peer","port":20857,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"117.8.100.81:24807"}
{"level":"info","ts":1583740765.141455,"msg":"announce_peer","addr":"114.26.40.207:64525","port":8740,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740766.8943725,"msg":"announce_peer","addr":"60.139.158.189:16892","port":16892,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740768.2217493,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"114.26.40.207:64525","port":8740}
{"level":"info","ts":1583740769.1918473,"msg":"announce_peer","addr":"114.26.40.207:64525","port":8740,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740769.2813182,"msg":"announce_peer","addr":"39.83.31.215:1024","port":13635,"infoHash":"9be16a021de1a764785fc0af1868115e918cbf61"}
{"level":"info","ts":1583740769.794455,"msg":"announce_peer","addr":"59.126.130.146:19985","port":19985,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740770.2917666,"msg":"announce_peer","addr":"114.26.40.207:64525","port":8740,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740770.860387,"msg":"announce_peer","addr":"222.248.16.131:24213","port":24213,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740771.0644333,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"220.228.196.10:1873","port":13421}
{"level":"info","ts":1583740771.1963315,"msg":"announce_peer","addr":"222.248.16.131:24213","port":24213,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740771.9917493,"msg":"announce_peer","addr":"114.26.40.207:64525","port":8740,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740772.869181,"msg":"announce_peer","addr":"39.83.31.215:1024","port":13635,"infoHash":"9be16a021de1a764785fc0af1868115e918cbf61"}
{"level":"info","ts":1583740773.7018223,"msg":"announce_peer","addr":"49.169.76.7:40703","port":40703,"infoHash":"9be16925eaa1518a8fd02a9ba32935312468507a"}
{"level":"info","ts":1583740775.0405757,"msg":"announce_peer","addr":"222.248.16.131:24213","port":24213,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740776.3463855,"msg":"announce_peer","addr":"114.26.40.207:64525","port":8740,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740776.8790677,"msg":"announce_peer","addr":"115.150.80.118:1040","port":25138,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740777.2070231,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"115.150.80.118:1040","port":25138}
{"level":"info","ts":1583740777.403968,"msg":"announce_peer","addr":"39.83.31.215:1024","port":13635,"infoHash":"9be16a021de1a764785fc0af1868115e918cbf61"}
{"level":"info","ts":1583740778.1561604,"msg":"announce_peer","addr":"111.201.229.31:4277","port":12755,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740778.3633702,"msg":"announce_peer","addr":"117.8.100.81:24807","port":20857,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740778.5801315,"msg":"announce_peer","port":45709,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"14.111.58.64:10773"}
{"level":"info","ts":1583740778.580323,"msg":"announce_peer","addr":"42.197.33.9:22223","port":22223,"infoHash":"99da003a28f1e0b869643ee71330ecdb8342b588"}
{"level":"info","ts":1583740778.9616086,"msg":"announce_peer","addr":"42.197.33.9:22223","port":22223,"infoHash":"99da003a28f1e0b869643ee71330ecdb8342b588"}
{"level":"info","ts":1583740779.2138188,"msg":"announce_peer","addr":"220.228.196.10:1873","port":13421,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740779.8654003,"msg":"announce_peer","port":24213,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"222.248.16.131:24213"}
{"level":"info","ts":1583740780.094627,"msg":"announce_peer","port":24213,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"222.248.16.131:24213"}
{"level":"info","ts":1583740780.984471,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"222.248.16.131:24213","port":24213}
{"level":"info","ts":1583740781.0882967,"msg":"announce_peer","addr":"43.250.108.207:24710","port":23060,"infoHash":"9bb6330b35e77fee9a6e4208da2f3c4b71904f93"}
{"level":"info","ts":1583740781.219295,"msg":"announce_peer","port":19985,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"59.126.130.146:19985"}
{"level":"info","ts":1583740781.436348,"msg":"announce_peer","addr":"212.124.7.87:9452","port":35663,"infoHash":"9be168339c4269a938ebfb2da1679257b5d48a52"}
{"level":"info","ts":1583740781.9694211,"msg":"announce_peer","addr":"220.228.196.10:1873","port":13421,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740782.0440805,"msg":"announce_peer","addr":"60.139.158.189:16892","port":16892,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740782.1014407,"msg":"announce_peer","infoHash":"9aa989a1d4c3685404adf51abf83fb723666db14","addr":"223.192.194.145:4364","port":9100}
{"level":"info","ts":1583740782.6560383,"msg":"announce_peer","port":20857,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"117.8.100.81:24807"}
{"level":"info","ts":1583740782.6691449,"msg":"announce_peer","addr":"117.8.100.81:24807","port":20857,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740782.719281,"msg":"announce_peer","addr":"118.150.254.126:42350","port":42350,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740783.4812863,"msg":"announce_peer","addr":"42.197.33.9:22223","port":22223,"infoHash":"99da003a28f1e0b869643ee71330ecdb8342b588"}
{"level":"info","ts":1583740783.7159548,"msg":"announce_peer","port":12755,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"111.201.229.31:4277"}
{"level":"info","ts":1583740784.6251774,"msg":"announce_peer","addr":"115.204.198.182:37896","port":37896,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740784.9777164,"msg":"announce_peer","port":20857,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"117.8.100.81:24807"}
{"level":"info","ts":1583740785.405472,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"111.182.81.4:16418","port":51392}
{"level":"info","ts":1583740786.5728052,"msg":"announce_peer","addr":"114.26.40.207:64525","port":8740,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740787.4742415,"msg":"announce_peer","addr":"117.8.100.81:24807","port":20857,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740787.7065392,"msg":"announce_peer","addr":"111.201.229.31:4277","port":12755,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740787.858501,"msg":"announce_peer","port":49161,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"222.143.22.194:15938"}
{"level":"info","ts":1583740788.1566432,"msg":"announce_peer","addr":"42.197.33.9:22223","port":22223,"infoHash":"99da003a28f1e0b869643ee71330ecdb8342b588"}
{"level":"info","ts":1583740789.7028086,"msg":"announce_peer","addr":"103.63.154.154:6718","port":22261,"infoHash":"9be59d184594795e62a2e09d164e7d796bbbcc2f"}
{"level":"info","ts":1583740789.7049983,"msg":"announce_peer","addr":"115.150.80.118:1040","port":25138,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740790.1650805,"msg":"announce_peer","addr":"115.150.80.118:1040","port":25138,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740790.5892825,"msg":"announce_peer","addr":"114.26.40.207:64525","port":8740,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740790.8257222,"msg":"announce_peer","addr":"60.139.158.189:16892","port":16892,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740792.2645924,"msg":"announce_peer","addr":"222.248.16.131:24213","port":24213,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740793.1378536,"msg":"announce_peer","addr":"117.8.100.81:24807","port":20857,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740793.4976876,"msg":"announce_peer","port":21487,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"124.218.43.159:21487"}
{"level":"info","ts":1583740794.5922437,"msg":"announce_peer","addr":"222.248.16.131:24213","port":24213,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740794.6535685,"msg":"announce_peer","addr":"113.71.213.98:21504","port":21390,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740794.6783698,"msg":"announce_peer","addr":"114.26.40.207:64525","port":8740,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740795.1863842,"msg":"announce_peer","addr":"180.90.250.69:7879","port":7879,"infoHash":"9bb0bc282b3cc79c01fa46f7a3bc778a61191252"}
{"level":"info","ts":1583740795.6946764,"msg":"announce_peer","port":24213,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"222.248.16.131:24213"}
{"level":"info","ts":1583740795.7615097,"msg":"announce_peer","addr":"113.71.213.98:21504","port":21390,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740796.1819694,"msg":"announce_peer","port":7879,"infoHash":"9be182179834ed07446311a6bf516dccec6c62f3","addr":"180.90.250.69:7879"}
{"level":"info","ts":1583740796.268276,"msg":"announce_peer","addr":"220.228.196.10:1873","port":13421,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740796.3770463,"msg":"announce_peer","port":22223,"infoHash":"99da003a28f1e0b869643ee71330ecdb8342b588","addr":"42.197.33.9:22223"}
{"level":"info","ts":1583740796.8755255,"msg":"announce_peer","infoHash":"99da003a28f1e0b869643ee71330ecdb8342b588","addr":"42.197.33.9:22223","port":22223}
{"level":"info","ts":1583740797.2739327,"msg":"announce_peer","port":20722,"infoHash":"9be168339c4269a938ebfb2da1679257b5d48a52","addr":"178.126.119.110:20722"}
{"level":"info","ts":1583740797.3582816,"msg":"announce_peer","addr":"222.248.16.131:24213","port":24213,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740797.867138,"msg":"announce_peer","port":62888,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"36.85.206.83:1161"}
{"level":"info","ts":1583740798.137312,"msg":"announce_peer","addr":"223.89.26.77:40305","port":8006,"infoHash":"9be53a8dd92bae355154748d2304b6ce250724ca"}
{"level":"info","ts":1583740799.30156,"msg":"announce_peer","addr":"115.190.38.240:30220","port":30220,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740801.4644449,"msg":"announce_peer","addr":"133.232.164.31:37246","port":37246,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740802.5984719,"msg":"announce_peer","infoHash":"9beff0acf5bc96a863c8b17fa3af2b8aa5de9448","addr":"103.27.25.174:33508","port":24158}
{"level":"info","ts":1583740803.2486782,"msg":"announce_peer","addr":"218.187.101.123:17724","port":17724,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740804.0876975,"msg":"announce_peer","addr":"42.197.33.9:22223","port":22223,"infoHash":"99da003a28f1e0b869643ee71330ecdb8342b588"}
{"level":"info","ts":1583740804.092976,"msg":"announce_peer","addr":"220.228.196.10:1873","port":13421,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740804.176038,"msg":"announce_peer","addr":"36.233.60.250:15107","port":15107,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740804.248828,"msg":"announce_peer","addr":"61.228.7.69:44496","port":44496,"infoHash":"9be164d0d0e02c73e1729adaab33b5846f87389f"}
{"level":"info","ts":1583740804.329885,"msg":"announce_peer","addr":"211.33.75.144:56262","port":56262,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740805.153661,"msg":"announce_peer","addr":"114.26.40.207:64525","port":8740,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740805.8065333,"msg":"announce_peer","addr":"112.14.72.226:6027","port":22353,"infoHash":"9be16a73baa0fad31731f5d57ddf5e61a9891a67"}
{"level":"info","ts":1583740806.0835092,"msg":"announce_peer","addr":"220.228.196.10:1873","port":13421,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740807.191883,"msg":"announce_peer","infoHash":"99da003a28f1e0b869643ee71330ecdb8342b588","addr":"42.197.33.9:22223","port":22223}
{"level":"info","ts":1583740807.404489,"msg":"announce_peer","port":24734,"infoHash":"9be168339c4269a938ebfb2da1679257b5d48a52","addr":"178.155.5.87:34495"}
{"level":"info","ts":1583740808.82767,"msg":"announce_peer","addr":"220.228.196.10:1873","port":13421,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740809.7548618,"msg":"announce_peer","addr":"14.131.205.174:1024","port":22943,"infoHash":"9bddadf03bcaa64cca4849889c2bc8deb43b89f4"}
{"level":"info","ts":1583740811.272396,"msg":"announce_peer","port":14396,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"122.42.158.38:14396"}
{"level":"info","ts":1583740811.8558154,"msg":"announce_peer","addr":"124.244.122.59:18864","port":18864,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740812.5025854,"msg":"announce_peer","port":13421,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"220.228.196.10:1873"}
{"level":"info","ts":1583740812.8451705,"msg":"announce_peer","addr":"43.227.138.69:30207","port":9304,"infoHash":"9be94f66cee56c3d63a32aa9554bb9ff4150cb1a"}
{"level":"info","ts":1583740813.9324887,"msg":"announce_peer","addr":"14.131.93.231:12951","port":12951,"infoHash":"9bd56482d6fd6a436f5051d3b9560cdd942a5962"}
{"level":"info","ts":1583740815.1942368,"msg":"announce_peer","addr":"220.228.196.10:1873","port":13421,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740815.644745,"msg":"announce_peer","addr":"220.228.196.10:1873","port":13421,"infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb"}
{"level":"info","ts":1583740816.4961476,"msg":"announce_peer","infoHash":"9be1701a5ac5a93f541b11a8e2fa24b562fa66bb","addr":"124.244.122.59:18864","port":18864}`
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
