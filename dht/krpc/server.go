package krpc

import (
	"github.com/GalaIO/P2Pcrawler/misc"
)

var supportQueryType = misc.List{"ping", "find_node", "get_peers", "announce_peer"}

var defaultConn = StartUp(":21000")
var defaultTxIdGen = NewTxIdGenerator(100)
var queriesHandlerMap = make(map[string]QueryHandler, 4)

// query handle context
type QueryCtx struct {
	txId     string
	qType    string
	sourceId string
	body     misc.Dict
}

func NewQueryCtx(txId, sourceId, qType string, body misc.Dict) *QueryCtx {
	return &QueryCtx{
		txId:     txId,
		sourceId: sourceId,
		qType:    qType,
		body:     body,
	}
}

type QueryHandler func(ctx *QueryCtx) misc.Dict

func RegisteQueryHandler(qType string, handler QueryHandler) {
	if !supportQueryType.ContainsString(qType) {
		panic("cannot support the query type")
	}

	if handler == nil {
		panic("register fail, handler is nil")
	}

	queriesHandlerMap[qType] = handler
}

//func BootStrap(host string) error {
//	defaultConn.SendPacketToHost(withFindNodeMsg())
//}

// query handler
func queryHandle(resp misc.Dict) (ret Response) {

	defer func() {
		if err := recover(); err != nil {
			dhtError, ok := err.(*misc.Error)
			if !ok {
				panic(err)
			}
			ret = withParamErr("aa", dhtError.Error())
		}
	}()

	// parse header
	txId := resp.GetString("t")
	queryType := resp.GetString("q")
	if !supportQueryType.ContainsString(queryType) {
		return withParamErr(txId, "donnot support <"+queryType+"> query type")
	}
	body := resp.GetDict("a")
	sourceId := body.GetString("id")
	if len(sourceId) != 20 {
		return withParamErr(txId, "id format err")
	}

	// do handle
	handler := queriesHandlerMap[queryType]
	return withResponse(txId, handler(NewQueryCtx(txId, sourceId, queryType, body)))
}
