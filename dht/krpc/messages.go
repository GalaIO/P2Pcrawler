package krpc

import "github.com/GalaIO/P2Pcrawler/dht"

// define query msg
func withPingMsg(txId string, nodeId string) dht.Dict {
	return withQueryMsg(txId, "ping", dht.Dict{
		"id": nodeId,
	})
}

func withFindNodeMsg(txId string, nodeId, target string) dht.Dict {

	return withQueryMsg(txId, "find_node", dht.Dict{
		"id":     nodeId,
		"target": target,
	})
}
