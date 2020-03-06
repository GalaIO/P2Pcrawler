package krpc

import (
	"github.com/GalaIO/P2Pcrawler/misc"
	"net"
)

var log = misc.GetLogger().SetPrefix("udp")

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
	log.Trace("conn close...", misc.Dict{"laddr": s.laddr})
	return s.conn.Close()
}

func (s *UdpServer) SendPacket(bytes []byte, raddr *net.UDPAddr) error {
	n, err := s.conn.WriteToUDP(bytes, raddr)
	if err != nil {
		log.Error(">>> send udp err", misc.Dict{"laddr": s.laddr, "err": err})
		return err
	}
	log.Info(">>> send Bytes", misc.Dict{"laddr": s.laddr, "raddr": raddr.String(), "length": n})
	return nil
}

func (s *UdpServer) SendPacketToHost(bytes []byte, remoteAddr string) error {
	raddr, err := net.ResolveUDPAddr("udp", remoteAddr)
	if err != nil {
		log.Error("resolve remote host err", misc.Dict{"raddr": remoteAddr, "err": err})
		return err
	}

	return s.SendPacket(bytes, raddr)
}

func (s *UdpServer) RecvChan() chan RecvPacket {
	return s.recvq
}

func StartUp(localAddr string) *UdpServer {
	laddr, err := net.ResolveUDPAddr("udp", localAddr)
	if err != nil {
		log.Panic("resolve local host err", misc.Dict{"laddr": laddr, "err": err})
	}

	serverConn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Panic("listen udp err", misc.Dict{"laddr": laddr, "err": err})
	}
	log.Trace("listen udp...", misc.Dict{"laddr": laddr})

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
			log.Error("<<< receive udp err", misc.Dict{"err": err})
			continue
		}
		log.Info("<<< received Bytes", misc.Dict{"raddr": raddr.String(), "length": n})
		srv.recvq <- RecvPacket{Addr: raddr, Bytes: buf[:n]}
	}
}
