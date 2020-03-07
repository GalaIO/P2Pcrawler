package dht

import (
	"github.com/GalaIO/P2Pcrawler/dht/krpc"
	"github.com/GalaIO/P2Pcrawler/misc"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

var testNodeId = "abcdefghij0123456789"

func TestGenQueryMsg(t *testing.T) {

	var req krpc.Request

	req = WithPingMsg("aa", testNodeId, nil)
	assert.Equal(t, misc.Dict{"t": "aa", "y": "q", "q": "ping", "a": misc.Dict{"id": "abcdefghij0123456789"}}, req.RawData())

	req = WithFindNodeMsg("aa", "abcdefghij0123456789", "mnopqrstuvwxyz123456", nil)
	assert.Equal(t, misc.Dict{"t": "aa", "y": "q", "q": "find_node", "a": misc.Dict{"id": "abcdefghij0123456789", "target": "mnopqrstuvwxyz123456"}}, req.RawData())

	req = WithGetPeersMsg("aa", "abcdefghij0123456789", "mnopqrstuvwxyz123456", nil)
	assert.Equal(t, misc.Dict{"t": "aa", "y": "q", "q": "get_peers", "a": misc.Dict{"id": "abcdefghij0123456789", "info_hash": "mnopqrstuvwxyz123456"}}, req.RawData())

	req = WithAnnouncePeerMsg("aa", "abcdefghij0123456789", "mnopqrstuvwxyz123456", "aoeusnth", 6881, nil)
	assert.Equal(t, misc.Dict{"t": "aa", "y": "q", "q": "announce_peer", "a": misc.Dict{"id": "abcdefghij0123456789", "implied_port": 0, "info_hash": "mnopqrstuvwxyz123456", "port": 6881, "token": "aoeusnth"}}, req.RawData())
}

func TestGenErrMsg(t *testing.T) {

	var err krpc.Response

	err = krpc.WithErr("aa", 202, "test protocol err")
	assert.Equal(t, misc.Dict{"t": "aa", "y": "e", "e": misc.List{202, "test protocol err"}}, err.RawData())

	err = krpc.WithParamErr("aa", "test param err")
	assert.Equal(t, misc.Dict{"t": "aa", "y": "e", "e": misc.List{203, "test param err"}}, err.RawData())
}

func TestGenRespMsg(t *testing.T) {

	var resp krpc.Response

	resp = WithPingResponse("aa", "mnopqrstuvwxyz123456")
	assert.Equal(t, misc.Dict{"t": "aa", "y": "r", "r": misc.Dict{"id": "mnopqrstuvwxyz123456"}}, resp.RawData())
}

func TestHandleQueryResp(t *testing.T) {

	var req krpc.Request

	req = WithPingMsg("aa", testNodeId, func(req krpc.Request, resp krpc.Response) {
		assert.Equal(t, "aa", resp.TxId())
		assert.Equal(t, false, resp.Error())
		assert.Equal(t, "mnopqrstuvwxyz123456", resp.NodeId())
		body := resp.Body()
		assert.Equal(t, "mnopqrstuvwxyz123456", body.GetString("id"))
		assert.Equal(t, misc.Dict{"id": "mnopqrstuvwxyz123456"}, body)
	})
	req.Handler()(req, WithPingResponse(req.TxId(), "mnopqrstuvwxyz123456"))

	req = WithFindNodeMsg("aa", "abcdefghij0123456789", "mnopqrstuvwxyz123456", func(req krpc.Request, resp krpc.Response) {
		body := resp.Body()
		nodes := body.GetString("nodes")
		nodeInfos := parseNodeInfo(nodes)
		assert.Equal(t, "mnopqrstuvwxyz123456", nodeInfos[0].Id)
		assert.Equal(t, "127.0.0.1:9000", nodeInfos[0].Addr.String())
	})
	req.Handler()(req, WithFindNodeResponse(req.TxId(), "mnopqrstuvwxyz123456", []*NodeInfo{NewNodeInfoFromHost("mnopqrstuvwxyz123456", "127.0.0.1:9000")}))

	req = WithGetPeersMsg("aa", "abcdefghij0123456789", "mnopqrstuvwxyz123456", func(req krpc.Request, resp krpc.Response) {
		body := resp.Body()
		assert.True(t, body.Exist("values"))
		vals := body.GetList("values")
		assert.Equal(t, "127.0.0.1:9000", parsePeerInfo(vals)[0].String())
	})
	req.Handler()(req, WithGetPeersValsResponse(req.TxId(), "mnopqrstuvwxyz123456", "aoeusnth", []*net.UDPAddr{resolveHost("127.0.0.1:9000")}))

	req = WithGetPeersMsg("aa", "abcdefghij0123456789", "mnopqrstuvwxyz123456", func(req krpc.Request, resp krpc.Response) {
		body := resp.Body()
		assert.True(t, body.Exist("nodes"))
		infos := body.GetString("nodes")
		assert.Equal(t, "127.0.0.1:9000", parseNodeInfo(infos)[0].Addr.String())
	})
	req.Handler()(req, WithGetPeersNodesResponse(req.TxId(), "mnopqrstuvwxyz123456", "aoeusnth", []*NodeInfo{NewNodeInfoFromHost("mnopqrstuvwxyz123456", "127.0.0.1:9000")}))

	req = WithAnnouncePeerMsg("aa", "abcdefghij0123456789", "mnopqrstuvwxyz123456", "aoeusnth", 6881, func(req krpc.Request, resp krpc.Response) {
		body := resp.Body()
		assert.Equal(t, "mnopqrstuvwxyz123456", body.GetString("id"))
	})
	req.Handler()(req, WithAnnouncePeerResponse(req.TxId(), "mnopqrstuvwxyz123456"))
}
