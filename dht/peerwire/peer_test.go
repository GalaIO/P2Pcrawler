package peerwire

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestExpire(t *testing.T) {
	peerConn := NewPeerConn("", nil, nil, 1*time.Second)
	assert.Equal(t, false, peerConn.IsReadLoopExpire())
	time.Sleep(1 * time.Second)
	assert.Equal(t, true, peerConn.IsReadLoopExpire())
}
