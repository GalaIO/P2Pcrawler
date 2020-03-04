package main

import (
	"fmt"
	"github.com/GalaIO/P2Pcrawler/dht"
	"io/ioutil"
	"net"
)

func main() {

	conn, err := net.Dial("udp", "dht.transmissionbt.com:6881")
	if err != nil {
		panic(err)
	}

	//bmsg, err := dht.EncodeDict(dht.Dict{"t":"aa", "y":"q", "q":"ping", "a": dht.Dict{"id":"abcdefghij0123456789"}})
	bmsg, err := dht.EncodeDict(dht.Dict{"t": "aa", "y": "q", "q": "find_node", "a": dht.Dict{"id": "abc11fghij0123456789", "target": "mnopqrstuvwxyz123456"}})
	if err != nil {
		panic(err)
	}
	_, err = conn.Write([]byte(bmsg))
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}

	resp, err := dht.DecodeDict(string(bytes))
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
