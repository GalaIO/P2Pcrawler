package test

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"testing"
)

func TestUdp(t *testing.T) {
	laddr := ":9000"
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
