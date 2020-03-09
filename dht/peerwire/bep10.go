package peerwire

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/GalaIO/P2Pcrawler/misc"
)

// refer http://www.bittorrent.org/beps/bep_0010.html
const handShakeMsgId = 0

type ExtendedHandShakeMsg interface {
	Metadata() misc.Dict
	Dict() misc.Dict
}

type BaseExHandShakeMsg struct {
	dict misc.Dict
}

func NewBaseExHandShakeMsg(metadata misc.Dict, dict misc.Dict) *BaseExHandShakeMsg {
	if dict == nil {
		dict = misc.Dict{}
	}
	dict["m"] = metadata
	return &BaseExHandShakeMsg{dict: dict}
}

func (b *BaseExHandShakeMsg) Metadata() misc.Dict {
	return b.dict.GetDict("m")
}

func (b *BaseExHandShakeMsg) Dict() misc.Dict {
	return b.dict
}

func (b *BaseExHandShakeMsg) Bytes() []byte {
	var buf bytes.Buffer
	dst, err := misc.EncodeDict(b.dict)
	if err != nil {
		peerWireLogger.Panic("encode extended handshake err", misc.Dict{"dict": b.dict, "err": err})
	}
	preLenBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(preLenBytes, uint32(len(dst)+2))
	buf.Write(preLenBytes)
	buf.WriteByte(byte(ExtendedPeerMsg)) // peer msg type, extended msg
	buf.WriteByte(byte(handShakeMsgId))  // msgId = 0 extended handshake msg
	buf.Write([]byte(dst))
	return buf.Bytes()
}

// extended handshake, without prefixlen
func parseExtendedHandShake(data []byte) ExtendedHandShakeMsg {
	fmt.Println("parseExtendedHandShake", hex.EncodeToString(data))
	msgType := PeerMsgType(data[0])
	msgId := int(data[1])

	if ExtendedPeerMsg != msgType || handShakeMsgId != msgId {
		peerWireLogger.Panic("extended handshake resp err", misc.Dict{"msgType": msgType, "msgId": msgId})
	}
	dict, err := misc.DecodeDict(string(data[2:]))
	if err != nil {
		peerWireLogger.Panic("decode extended handshake err", misc.Dict{"err": err})
	}
	return NewBaseExHandShakeMsg(dict.GetDict("m"), dict)
}

func withExtendedhandShake(metadata misc.Dict, dict misc.Dict) []byte {
	exMsg := NewBaseExHandShakeMsg(metadata, dict)
	return exMsg.Bytes()
}
