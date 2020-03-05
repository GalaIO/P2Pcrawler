package test

import (
	"encoding/hex"
	"fmt"
	"github.com/GalaIO/P2Pcrawler/dht"
	"io/ioutil"
	"net"
	"os"
	"testing"
)

func TestPingNode(t *testing.T) {

	//conn, err := net.Dial("udp", "dht.transmissionbt.com:6881")
	conn, err := net.Dial("udp", "87.98.162.88:6881")
	if err != nil {
		t.Error(err)
	}

	//bmsg, err := dht.EncodeDict(dht.Dict{"t":"aa", "y":"q", "q":"ping", "a": dht.Dict{"id":"abcdefghij0123456789"}})
	//bmsg, err := dht.EncodeDict(dht.Dict{"t": "aa", "y": "q", "q": "find_node", "a": dht.Dict{"id": "abc11fghij0123456789", "target": "mnopqrstuvwxyz123456"}})
	//if err != nil {
	//	t.Error(err)
	//}
	_, err = conn.Write([]byte("d1:ad2:id20:abcdefghij01234567896:target20:mnopqrstuvwxyz123456e1:q9:find_node1:t2:aa1:y1:qe\n"))
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

func TestUdpTest(t *testing.T) {

	request := "d1:ad2:id20:abcdefghij01234567896:target20:mnopqrstuvwxyz123456e1:q9:find_node1:t2:aa1:y1:qe"
	rawAddr := "router.bittorrent.com:6881"
	//rawAddr:= "dht.transmissionbt.com:6881"
	//request := ""
	//rawAddr:= "time.nist.gov:37"
	laddr, err := net.ResolveUDPAddr("udp", ":6883")
	raddr, err := net.ResolveUDPAddr("udp", rawAddr)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		fmt.Println("Can't dial: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Println("failed:", err)
		os.Exit(1)
	}
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err)
		os.Exit(1)
	}
	fmt.Println(hex.EncodeToString(data[:n]))
	ret := string(data[:n])
	ds, err := dht.DecodeDict(ret)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body := ds.GetDict("r")
	fmt.Println(hex.EncodeToString([]byte(body.GetString("id"))))
	fmt.Println(hex.EncodeToString([]byte(body.GetString("nodes"))))
	fmt.Println(ret)

}
