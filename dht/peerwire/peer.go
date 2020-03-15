package peerwire

import (
	"crypto/sha1"
	"github.com/GalaIO/P2Pcrawler/misc"
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
	timeoutChan       <-chan time.Time
}

var peerBytesPool = misc.NewBytesPool(10000, 4)

func NewPeerConn(addr string, conn net.Conn, infoHash []byte, expireTime time.Duration) *PeerConn {
	return &PeerConn{
		addr:              addr,
		conn:              conn,
		infoHash:          infoHash,
		metaData:          make([][]byte, 16),
		requestedMetaData: []bool{false},
		connTimeStamp:     time.Now(),
		timeoutChan:       time.After(expireTime),
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
	select {
	case <-p.timeoutChan:
		return true
	default:
		return false
	}
}

func GeneratePeerId(str string) []byte {
	sha160 := sha1.Sum([]byte(str))
	return sha160[:]
}

func GenerateInfoHash(bytes []byte) []byte {
	sha160 := sha1.Sum(bytes)
	return sha160[:]
}
