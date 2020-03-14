package misc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"syscall"
	"testing"
	"time"
)

func TestShutdownByIntSig(t *testing.T) {
	ctx := WithShutdownContext(context.Background())
	ctx.sigChan <- syscall.SIGINT
	ctx.WaitShutdown(1 * time.Second)
	assert.Equal(t, Shutdowned, ctx.Status())
}

func TestShutdownByCtxCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	shutdownCtx := WithShutdownContext(ctx)
	shutdownCtx.WaitShutdown(1 * time.Second)
	assert.Equal(t, Shutdowned, shutdownCtx.Status())
}

func TestShutdownRunExitCallback(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	shutdownCtx := WithShutdownContext(ctx)
	count := 0
	shutdownCtx.Clean(func() {
		count++
		assert.Equal(t, Shutdowning, shutdownCtx.Status())
	}, func() {
		count++
	}, func() {
		time.Sleep(2 * time.Second)
		count++
	})
	shutdownCtx.WaitShutdown(1 * time.Second)
	assert.Equal(t, Shutdowned, shutdownCtx.Status())
	assert.Equal(t, 2, count)
}
