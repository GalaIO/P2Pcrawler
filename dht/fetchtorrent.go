package dht

import (
	"encoding/hex"
	"github.com/GalaIO/P2Pcrawler/dht/krpc"
	"github.com/GalaIO/P2Pcrawler/dht/peerwire"
	"github.com/GalaIO/P2Pcrawler/misc"
	"os"
)

var recvInfoHash = make(chan *krpc.RpcContext, 1000)

func fetchTorrent() {
	for {
		ctx := <-recvInfoHash
		fetchHandler(ctx)
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
	//impliedPort := body.GetInteger("implied_port")
	//laddr := ctx.RemoteAddr().String()
	//if impliedPort > 0 {
	//	laddr = ctx.RemoteAddr().String()
	//}
	hash := hex.EncodeToString([]byte(infoHash))
	go func() {
		result, err := peerwire.FetchMetaData(ctx.RemoteAddr().String(), peerwire.LocalPeerId, misc.Str2Bytes(infoHash))
		if err != nil {
			flushTorrentFile(hash, result)
		}
	}()
}

func flushTorrentFile(fileName string, data []byte) {
	file, err := os.Create("torrent/" + fileName + ".torrent")
	if err != nil {
		dhtLogger.Error("create file err", misc.Dict{"err": err})
		return
	}
	_, err = file.Write(data)
	if err != nil {
		dhtLogger.Error("write file err", misc.Dict{"err": err})
		return
	}
	file.Close()
	return
}
