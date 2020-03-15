package misc

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestBytesPool(t *testing.T) {
	bytesPool := NewBytesPool(100, 10)
	bytesOf10 := bytesPool.Get()
	bytesPool.Put(bytesOf10)
	bytes2Of10 := bytesPool.Get()
	assert.Equal(t, &bytesOf10, &bytes2Of10)
	bytesOf100 := bytesPool.GetBySize(100)
	bytesPool.Put(bytesOf100)
	bytes2Of100 := bytesPool.GetBySize(100)
	assert.Equal(t, &bytesOf100, &bytes2Of100)
}

func BenchmarkBytesPool(b *testing.B) {
	bytesPool := NewBytesPool(100, 10)
	for i := 0; i < b.N; i++ {
		length := ((i % 100) + 1) * 100000
		bytes := bytesPool.GetBySize(length)
		if len(bytes) != length {
			panic("get bytes err")
		}
		bytesPool.Put(bytes[:1])
	}
}

func BenchmarkBytesMalloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes := make([]byte, ((i%100)+1)*100000)
		if len(bytes) <= 0 {
			panic("get bytes err")
		}
	}
}
