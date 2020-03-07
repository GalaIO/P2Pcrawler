package krpc

import (
	"encoding/binary"
	"github.com/GalaIO/P2Pcrawler/misc"
	"net"
	"strings"
)

type NodeInfo struct {
	Id   string
	Addr *net.UDPAddr
}

func NewNodeInfo(id string, addr *net.UDPAddr) *NodeInfo {
	return &NodeInfo{Id: id, Addr: addr}
}

func NewNodeInfoFromHost(id, addr string) *NodeInfo {
	return &NodeInfo{Id: id, Addr: resolveHost(addr)}
}

func resolveHost(addr string) *net.UDPAddr {
	raddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		log.Panic("resolve remote host err", misc.Dict{"raddr": addr, "err": err})
	}
	return raddr
}

func joinNodeInfos(nodes []*NodeInfo) string {
	var builder strings.Builder
	builder.Grow(len(nodes) * 26)
	for _, v := range nodes {
		builder.WriteString(v.Id)
		builder.Write(v.Addr.IP.To4())
		portBytes := make([]byte, 2)
		binary.BigEndian.PutUint16(portBytes, uint16(v.Addr.Port))
		builder.Write(portBytes)
	}
	return builder.String()
}

func parseNodeInfo(infos string) []*NodeInfo {
	nodes := make([]*NodeInfo, len(infos)/26)
	for i := range nodes {
		j := i * 26
		bytes := []byte(infos[j+20 : j+26])
		nodes[i] = NewNodeInfo(infos[j:j+20], &net.UDPAddr{
			IP:   bytes[:4],
			Port: int(binary.BigEndian.Uint16(bytes[4:6])),
		})
	}
	return nodes
}

func parsePeerInfo(vals misc.List) []*net.UDPAddr {
	addrs := make([]*net.UDPAddr, len(vals))
	for i := range vals {
		val := vals.GetString(i)
		bytes := []byte(val)
		addrs[i] = &net.UDPAddr{
			IP:   bytes[:4],
			Port: int(binary.BigEndian.Uint16(bytes[4:6])),
		}
	}

	return addrs
}

func joinPeerInfos(addrs []*net.UDPAddr) misc.List {
	vals := make([]interface{}, len(addrs))
	for i, addr := range addrs {
		var builder strings.Builder
		builder.Write(addr.IP.To4())
		portBytes := make([]byte, 2)
		binary.BigEndian.PutUint16(portBytes, uint16(addr.Port))
		builder.Write(portBytes)
		vals[i] = builder.String()
	}
	return vals
}
