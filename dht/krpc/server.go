package krpc

import "github.com/GalaIO/P2Pcrawler/dht"

var supportQueryType dht.List = dht.List{"ping", "find_node", "get_peers", "announce_peer"}

// query handle context
type QueryCtx struct {
	txId     string
	qType    string
	sourceId string
	body     dht.Dict
}

func NewQueryCtx(txId, sourceId, qType string, body dht.Dict) *QueryCtx {
	return &QueryCtx{
		txId:     txId,
		sourceId: sourceId,
		qType:    qType,
		body:     body,
	}
}

type QueryHandler func(ctx *QueryCtx) dht.Dict

var queriesHandlerMap = make(map[string]QueryHandler, 4)

func RegisteQueryHandler(qType string, handler QueryHandler) {
	if !supportQueryType.ContainsString(qType) {
		panic("cannot support the query type")
	}

	if handler == nil {
		panic("register fail, handler is nil")
	}

	queriesHandlerMap[qType] = handler
}

// query handler
func queriesHandle(resp dht.Dict) (ret dht.Dict) {

	defer func() {
		if err := recover(); err != nil {
			dhtError, ok := err.(*dht.DhtError)
			if !ok {
				panic(err)
			}
			ret = withParamErr(dhtError.Error())
		}
	}()

	// parse header
	txId := resp.GetString("t")
	queryType := resp.GetString("q")
	if !supportQueryType.ContainsString(queryType) {
		return withParamErr("donnot support <" + queryType + "> query type")
	}
	body := resp.GetDict("a")
	sourceId := body.GetString("id")
	if len(sourceId) != 20 {
		return withParamErr("id format err")
	}

	// do handle
	handler := queriesHandlerMap[queryType]
	return withResponse(txId, handler(NewQueryCtx(txId, sourceId, queryType, body)))
}
