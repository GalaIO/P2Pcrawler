package misc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewWorkPool(t *testing.T) {
	pool := NewWorkPool(context.Background(), "testpool", 2)
	count := 0
	pool.AsyncSubmit(func() {
		count++
	})
	pool.AsyncSubmit(func() {
		count++
	})
	pool.Shutdown(1 * time.Second)
	assert.Equal(t, 2, count)
}

func TestNewWorkPoolTimeout(t *testing.T) {
	pool := NewWorkPool(context.Background(), "testpool", 2)
	count := 0
	pool.AsyncSubmit(func() {
		count++
	})
	pool.AsyncSubmit(func() {
		count++
	})
	pool.AsyncSubmit(func() {
		time.Sleep(2 * time.Second)
		count++
	})
	pool.Shutdown(1 * time.Second)
	assert.Equal(t, 2, count)
}

func TestSubmitCloseWorkPoolPanic(t *testing.T) {
	pool := NewWorkPool(context.Background(), "testpool", 2)
	count := 0
	defer func() {
		assert.Equal(t, 1, count)
		if err := recover(); err != nil {
			assert.Equal(t, "testpool is shutdowned", err)
		}
	}()
	pool.AsyncSubmit(func() {
		count++
	})
	pool.Shutdown(1 * time.Second)
	pool.AsyncSubmit(func() {
		count++
	})
}

func TestCannelWorkPool(t *testing.T) {
	ctx, cannel := context.WithCancel(context.Background())
	pool := NewWorkPool(ctx, "testpool", 2)
	count := 0
	defer func() {
		assert.Equal(t, 1, count)
		if err := recover(); err != nil {
			assert.Equal(t, "testpool is shutdowned", err)
		}
	}()
	pool.AsyncSubmit(func() {
		count++
	})
	// wait for execute
	time.Sleep(200 * time.Millisecond)
	cannel()
	time.Sleep(200 * time.Millisecond)
	pool.AsyncSubmit(func() {
		count++
	})
	time.Sleep(200 * time.Millisecond)
}
