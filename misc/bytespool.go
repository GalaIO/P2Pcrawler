package misc

import "sync"

// BytesPool implements a leaky pool of []byte in the form of a bounded
// channel.
type BytesPool struct {
	sync.RWMutex
	chanMapping  *sync.Map
	defaultWidth int
	maxSize      int
}

// NewBytesPool creates a new BytesPool bounded to the given maxSize, with new
// byte arrays sized based on width.
func NewBytesPool(maxSize int, width int) (bp *BytesPool) {
	chanMapping := new(sync.Map)
	chanMapping.Store(width, make(chan []byte, maxSize))
	return &BytesPool{
		chanMapping:  chanMapping,
		defaultWidth: width,
		maxSize:      maxSize,
	}
}

// Get gets a []byte from the BytesPool, or creates a new one if none are
// available in the pool.
func (bp *BytesPool) Get() (b []byte) {
	select {
	case b = <-bp.loadOrStore(bp.defaultWidth):
	// reuse existing buffer
	default:
		// create new buffer
		b = make([]byte, bp.defaultWidth)
	}
	return
}

func (bp *BytesPool) GetBySize(size int) (b []byte) {
	select {
	case b = <-bp.loadOrStore(size):
	// reuse existing buffer
	default:
		// create new buffer
		b = make([]byte, size)
	}
	return
}

// Put returns the given Buffer to the BytesPool.
// when chan is full, the bytes will ignore and collection when gc
func (bp *BytesPool) Put(b []byte) {
	bytesChan := bp.loadOrStore(cap(b))
	select {
	// reset bytes length, import...
	case bytesChan <- b[:cap(b)]:
		// buffer went back into pool
	default:
		// buffer didn't go back into pool, just discard
	}
}

// double check if exist chan, else create it.
func (bp *BytesPool) loadOrStore(i int) chan []byte {
	bp.chanMapping.Load(i)
	if c, exist := bp.chanMapping.Load(i); exist {
		return c.(chan []byte)
	}

	c, _ := bp.chanMapping.LoadOrStore(i, make(chan []byte, bp.maxSize))
	return c.(chan []byte)
}
