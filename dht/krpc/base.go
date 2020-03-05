package krpc

import "github.com/GalaIO/P2Pcrawler/dht"

// query
func withQueryMsg(txId string, queryType string, body dht.Dict) dht.Dict {
	return dht.Dict{
		"t": txId,
		"y": "q",
		"q": queryType,
		"a": body,
	}
}

// response
func withResponse(txId string, resp dht.Dict) dht.Dict {
	return dht.Dict{
		"t": txId,
		"y": "r",
		"r": resp,
	}
}

// error
func withParamErr(msg string) dht.Dict {
	if len(msg) <= 0 {
		msg = "invalid param"
	}
	return withErr(dht.ProtocolErr, msg)
}

func withErr(code dht.DhtErrCode, errMsg string) dht.Dict {
	return dht.Dict{"t": "aa", "y": "e", "e": dht.List{code, errMsg}}
}
