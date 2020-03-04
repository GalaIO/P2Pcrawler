package dht

func withPingMsg(txId string, nodeId string) Dict {
	return withQueryMsg(txId, "ping", Dict{
		"id": nodeId,
	})
}

func pingHandle(msg Dict, handleFunc func() string) (ret Dict) {

	defer func() {
		if err := recover(); err != nil {
			dhtError := err.(*DhtError)
			ret = withParamErr(dhtError.Error())
		}
	}()

	txId := msg.GetString("t")
	body := msg.GetDict("a")
	sourceId := body.GetString("id")

	if len(sourceId) != 20 {
		return withParamErr("id format err")
	}
	return withResponse(txId, Dict{
		"id": handleFunc(),
	})
}

func findNodeHandle(msg Dict, handleFunc func(nodeId string) (string, string)) (ret Dict) {

	defer func() {
		if err := recover(); err != nil {
			dhtError := err.(*DhtError)
			ret = withParamErr(dhtError.Error())
		}
	}()

	txId := msg.GetString("t")
	body := msg.GetDict("a")
	sourceId := body.GetString("id")

	if len(sourceId) != 20 {
		return withParamErr("id format err")
	}

	nodeId, nodes := handleFunc(body.GetString("target"))
	return withResponse(txId, Dict{
		"id":    nodeId,
		"nodes": nodes,
	})
}

func withFindNodeMsg(txId string, nodeId, target string) Dict {

	return withQueryMsg(txId, "find_node", Dict{
		"id":     nodeId,
		"target": target,
	})
}
func withQueryMsg(txId string, queryType string, body Dict) Dict {
	return Dict{
		"t": txId,
		"y": "q",
		"q": queryType,
		"a": body,
	}
}
func withResponse(txId string, resp Dict) Dict {
	return Dict{
		"t": txId,
		"y": "r",
		"r": resp,
	}
}
