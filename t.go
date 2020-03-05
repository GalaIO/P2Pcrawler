package main

import (
	"encoding/hex"
	"fmt"
	"github.com/GalaIO/P2Pcrawler/dht"
	"net"
	"os"
)

func main() {

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
