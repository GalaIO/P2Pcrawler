package krpc

import (
	"github.com/GalaIO/P2Pcrawler/misc"
	"log"
	"net"
)

var udpLogger = misc.GetLogger().SetPrefix("udp")

// define udp receive length, when packet is too large, will be ignore
var RecvPacketLen = 1024
var RecvChanLen = 100

type RecvPacket struct {
	Addr  *net.UDPAddr
	Bytes []byte
}

type UdpServer struct {
	laddr *net.UDPAddr
	conn  *net.UDPConn
	recvq chan RecvPacket
}

func (s *UdpServer) Close() error {
	udpLogger.Trace("conn close...", misc.Dict{"laddr": s.laddr})
	return s.conn.Close()
}

func (s *UdpServer) SendPacket(bytes []byte, raddr *net.UDPAddr) error {
	n, err := s.conn.WriteToUDP(bytes, raddr)
	if err != nil {
		udpLogger.Error(">>> send udp err", misc.Dict{"laddr": s.laddr, "err": err})
		return err
	}
	udpLogger.Info(">>> send Bytes", misc.Dict{"laddr": s.laddr, "raddr": raddr.String(), "length": n})
	return nil
}

func (s *UdpServer) SendPacketToHost(bytes []byte, remoteAddr string) error {
	raddr, err := net.ResolveUDPAddr("udp", remoteAddr)
	if err != nil {
		udpLogger.Error("resolve remote host err", misc.Dict{"raddr": remoteAddr, "err": err})
		return err
	}

	return s.SendPacket(bytes, raddr)
}

func (s *UdpServer) RecvChan() chan RecvPacket {
	return s.recvq
}

// startup a udp server, listening on target port
func StartUp(localAddr string) *UdpServer {
	laddr, err := net.ResolveUDPAddr("udp", localAddr)
	if err != nil {
		udpLogger.Panic("resolve local host err", misc.Dict{"laddr": laddr, "err": err})
	}

	serverConn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Panic("listen udp err", misc.Dict{"laddr": laddr, "err": err})
	}
	udpLogger.Trace("listen udp...", misc.Dict{"laddr": laddr})

	server := &UdpServer{
		laddr: laddr,
		conn:  serverConn,
		recvq: make(chan RecvPacket, RecvChanLen),
	}
	go recvRoutiue(server)
	return server
}

func recvRoutiue(srv *UdpServer) {
	conn := srv.conn
	for {
		buf := make([]byte, RecvPacketLen)
		n, raddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			udpLogger.Error("<<< receive udp err", misc.Dict{"err": err})
			continue
		}
		udpLogger.Info("<<< received Bytes", misc.Dict{"raddr": raddr.String(), "length": n})
		srv.recvq <- RecvPacket{Addr: raddr, Bytes: buf[:n]}
	}
}
