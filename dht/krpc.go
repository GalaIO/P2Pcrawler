package dht

func withPingMsg(txId string, nodeId string) Dict {
	return Dict{
		"t": txId,
		"y": "q",
		"q": "ping",
		"a": Dict{
			"id": nodeId,
		},
	}
}

func pingHandle(msg Dict, handleFunc func() string) Dict {

	body, exist := msg["a"]
	if !exist {
		return withParamErr("")
	}

	cmap, ok := body.(Dict)
	if !ok {
		return withParamErr("")
	}

	sourceId, exist := cmap["id"]
	if !exist {
		return withParamErr("id cannot be empty")
	}

	id, ok := sourceId.(string)
	if !ok || len(id) != 20 {
		return withParamErr("id format err")
	}

	return Dict{
		"t": msg["t"],
		"y": "r",
		"r": Dict{
			"id": handleFunc(),
		},
	}
}

func findNodeHandle(msg Dict, handleFunc func(nodeId string) (string, string)) Dict {
	body, _ := msg["a"]
	cmap, _ := body.(Dict)
	nodeId, nodes := handleFunc(cmap["target"].(string))
	return Dict{
		"t": msg["t"],
		"y": "r",
		"r": Dict{
			"id":    nodeId,
			"nodes": nodes,
		},
	}
}

func withFindNodeMsg(txId string, nodeId, target string) Dict {
	return Dict{
		"t": txId,
		"y": "q",
		"q": "find_node",
		"a": Dict{
			"id":     nodeId,
			"target": target,
		},
	}
}
