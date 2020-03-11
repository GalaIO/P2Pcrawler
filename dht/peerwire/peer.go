package peerwire

import (
	"crypto/sha1"
	"net"
	"time"
)

type PeerConn struct {
	addr              string
	conn              net.Conn
	infoHash          []byte
	metaData          [][]byte
	metaDataSize      int
	requestedMetaData []bool
	connTimeStamp     time.Time
	expireTime        time.Duration
}

func NewPeerConn(addr string, conn net.Conn, infoHash []byte, expireTime time.Duration) *PeerConn {
	return &PeerConn{
		addr:              addr,
		conn:              conn,
		infoHash:          infoHash,
		metaData:          make([][]byte, 16),
		requestedMetaData: []bool{false},
		connTimeStamp:     time.Now(),
		expireTime:        expireTime,
	}
}

func (p *PeerConn) CompledMetaData() bool {
	for _, v := range p.requestedMetaData {
		if !v {
			return false
		}
	}
	return true
}

func (p *PeerConn) IsReadLoopExpire() bool {
	if time.Now().After(p.connTimeStamp.Add(p.expireTime)) {
		return true
	}
	return false
}

func GeneratePeerId(str string) []byte {
	sha160 := sha1.Sum([]byte(str))
	return sha160[:]
}

func GenerateInfoHash(bytes []byte) []byte {
	sha160 := sha1.Sum(bytes)
	return sha160[:]
}
