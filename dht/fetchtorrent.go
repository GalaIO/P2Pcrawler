package dht

import (
	"encoding/hex"
	"github.com/GalaIO/P2Pcrawler/dht/krpc"
	"github.com/GalaIO/P2Pcrawler/dht/peerwire"
	"github.com/GalaIO/P2Pcrawler/misc"
	"os"
	"strconv"
)

var recvInfoHash = make(chan *krpc.RpcContext, 300000)

func fetchTorrent() {
	// 开启多个线程消费
	for i := 0; i < 10; i++ {
		go func() {
			for {
				ctx := <-recvInfoHash
				fetchHandler(ctx)
			}
		}()
	}
}

func fetchHandler(ctx *krpc.RpcContext) {

	defer func() {
		if err := recover(); err != nil {
			dhtLogger.Error("fetchHandler panic", misc.Dict{"err": err})
		}
	}()
	req := ctx.Request()
	body := req.Body()
	infoHash := body.GetString("info_hash")
	laddr := parseFetchAddr(ctx)
	hash := hex.EncodeToString([]byte(infoHash))
	result, err := peerwire.FetchMetaData(laddr, peerwire.LocalPeerId, misc.Str2Bytes(infoHash))
	if err != nil {
		return
	}
	flushTorrentFile(hash, result)
}

func parseFetchAddr(ctx *krpc.RpcContext) string {
	req := ctx.Request()
	body := req.Body()
	port := body.GetInteger("port")
	useUdpPort := false
	if body.Exist("implied_port") && body.GetInteger("implied_port") > 0 {
		useUdpPort = true
	}
	laddr := ctx.RemoteAddr().IP.String() + ":" + strconv.Itoa(port)
	if useUdpPort {
		laddr = ctx.RemoteAddr().String()
	}
	return laddr
}

func flushTorrentFile(fileName string, data []byte) {
	file, err := os.Create("torrent/" + fileName + ".torrent")
	defer file.Close()
	if err != nil {
		dhtLogger.Error("create file err", misc.Dict{"err": err})
		return
	}
	_, err = file.Write(data)
	if err != nil {
		dhtLogger.Error("write file err", misc.Dict{"err": err})
		return
	}
}
