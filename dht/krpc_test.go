package dht

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testNodeId = "abcdefghij0123456789"

func TestPingMsg(t *testing.T) {
	msg := withPingMsg("aa", testNodeId)
	assert.Equal(t, Dict{"t": "aa", "y": "q", "q": "ping", "a": Dict{"id": "abcdefghij0123456789"}}, msg)
	resp := pingHandle(msg, func() string {
		return "mnopqrstuvwxyz123456"
	})
	assert.Equal(t, Dict{"t": "aa", "y": "r", "r": Dict{"id": "mnopqrstuvwxyz123456"}}, resp)

	resp = pingHandle(Dict{"t": "aa", "y": "q", "q": "ping"}, func() string {
		return "mnopqrstuvwxyz123456"
	})
	assert.Equal(t, Dict{"t": "aa", "y": "e", "e": List{ProtocolErr, "invalid param"}}, resp)
}

func TestFindNodeMsg(t *testing.T) {
	msg := withFindNodeMsg("aa", "abcdefghij0123456789", "mnopqrstuvwxyz123456")
	assert.Equal(t, Dict{"t": "aa", "y": "q", "q": "find_node", "a": Dict{"id": "abcdefghij0123456789", "target": "mnopqrstuvwxyz123456"}},
		msg)

	resp := findNodeHandle(msg, func(nodeId string) (string, string) {
		return "0123456789abcdefghij", "def456..."
	})
	assert.Equal(t, Dict{"t": "aa", "y": "r", "r": Dict{"id": "0123456789abcdefghij", "nodes": "def456..."}},
		resp)
}
