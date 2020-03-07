package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"reflect"
	"time"
)

func panicTest() (ret string) {

	defer func() {
		if err := recover(); err != nil {
			ret = err.(string)
		}
	}()

	panic("123")
	return "456"
}

func main() {
	fmt.Println("hello world")

	//miscTest()
	udpTest()
}

func udpTest() {
	laddr := ":9090"
	raddr := "87.98.162.88:6881"
	request := []byte("d1:ad2:id20:abcdefghij01234567896:target20:mnopqrstuvwxyz123456e1:q9:find_node1:t2:aa1:y1:qe")
	//raddr := "time.nist.gov:37"
	//request := []byte("")
	// Resolving Address
	localAddr, err := net.ResolveUDPAddr("udp", laddr)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	remoteAddr, err := net.ResolveUDPAddr("udp", raddr)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	// Build listening connections
	conn, err := net.ListenUDP("udp", localAddr)
	// Exit if some error occured
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	defer conn.Close()

	// write a message to server
	_, err = conn.WriteToUDP(request, remoteAddr)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(">>> Packet sent to: ", raddr)
	}

	// Receive response from server
	buf := make([]byte, 1024)
	rn, remAddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("<<<  %d bytes received from: %v, data: %s\n", rn, remAddr, hex.EncodeToString(buf[:rn]))
	}
}

func miscTest() {
	var l []interface{}
	l = nil
	for e := range l {
		fmt.Println(e)
	}
	var m map[string]interface{}
	m = nil
	for k, v := range m {
		fmt.Println(k, v)
	}
	s1 := sha1.New()
	s1.Write([]byte("test"))
	sha1sum := s1.Sum(nil)
	fmt.Println(hex.EncodeToString(sha1sum[:]))
	type Dict map[string]interface{}
	d := Dict{"name": "xiaoguo"}
	m1 := map[string]interface{}{"name": "xiaoguo"}
	t := reflect.TypeOf(d)
	mt := reflect.TypeOf(m1)
	fmt.Printf("%s\n", t.Name())
	fmt.Printf("%s\n", t.Kind())
	fmt.Printf("%s\n", mt.Name())
	fmt.Printf("%s\n", mt.Kind())
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))
	fmt.Println(panicTest())
}
