package krpc

import (
	"encoding/hex"
	"net"
	"testing"
)

func TestUdpServer_SendPacket(t *testing.T) {
	laddr := ":9000"
	raddr := "87.98.162.88:6881"
	request := []byte("d1:ad2:id20:abcdefghij01234567896:target20:mnopqrstuvwxyz123456e1:q9:find_node1:t2:aa1:y1:qe")
	//raddr := "time.nist.gov:37"
	//request := []byte("")

	udpServer := StartUp(laddr)

	remoteAddr, err := net.ResolveUDPAddr("udp", raddr)
	if err != nil {
		t.Fatal(err)
	}
	err = udpServer.SendPacket(request, remoteAddr)
	if err != nil {
		t.Fatal(err)
	}

	packet := <-udpServer.RecvChan()

	t.Logf("<<<  %d Bytes received from: %v, data: %s\n", len(packet.Bytes), packet.Addr, hex.EncodeToString(packet.Bytes))
}

func TestUdpServer_Listen(t *testing.T) {
	laddr := "localhost:9000"
	raddr := "localhost:6881"

	udpServer := StartUp(raddr)
	go func() {
		packet := <-udpServer.RecvChan()
		t.Logf("[server]<<<  %d Bytes received from: %v, data: %s\n", len(packet.Bytes), packet.Addr, hex.EncodeToString(packet.Bytes))

		err := udpServer.SendPacket(packet.Bytes, packet.Addr)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("[server]>>>  %d Bytes send to: %v, data: %s\n", len(packet.Bytes), packet.Addr, hex.EncodeToString(packet.Bytes))
	}()

	udpClient := StartUp(laddr)
	request := []byte("hello world")
	err := udpClient.SendPacketToHost(request, raddr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("[client]>>>  %d Bytes send to: %v, data: %s\n", len(request), raddr, hex.EncodeToString(request))
	packet := <-udpClient.RecvChan()
	t.Logf("[client]<<<  %d Bytes received from: %v, data: %s\n", len(packet.Bytes), packet.Addr, hex.EncodeToString(packet.Bytes))
}
