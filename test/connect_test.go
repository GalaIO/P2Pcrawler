package test

import (
	"github.com/GalaIO/P2Pcrawler/dht"
	"io/ioutil"
	"net"
	"testing"
)

func TestPingNode(t *testing.T) {

	conn, err := net.Dial("udp", "dht.transmissionbt.com:6881")
	if err != nil {
		t.Error(err)
	}

	//bmsg, err := dht.EncodeDict(dht.Dict{"t":"aa", "y":"q", "q":"ping", "a": dht.Dict{"id":"abcdefghij0123456789"}})
	bmsg, err := dht.EncodeDict(dht.Dict{"t": "aa", "y": "q", "q": "find_node", "a": dht.Dict{"id": "abc11fghij0123456789", "target": "mnopqrstuvwxyz123456"}})
	if err != nil {
		t.Error(err)
	}
	_, err = conn.Write([]byte(bmsg))
	if err != nil {
		t.Error(err)
	}

	bytes, err := ioutil.ReadAll(conn)
	if err != nil {
		t.Error(err)
	}

	resp, err := dht.DecodeDict(string(bytes))
	if err != nil {
		t.Error(err)
	}

	t.Log(resp)
}
