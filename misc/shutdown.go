package misc

import (
	"context"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type ShutdownStatus int

const (
	Running ShutdownStatus = iota
	Shutdowning
	Shutdowned
)

var shutdownLogger = GetLogger().SetPrefix("shutdown")

var defaultShutdownCtx = WithShutdownContext(context.Background())

func InitShutdownCtx(ctx context.Context) {
	defaultShutdownCtx.innerCtx = ctx
}

func RegisterShutDownClean(cleanFunc ...CleanFunc) {
	defaultShutdownCtx.Clean(cleanFunc...)
}

func Wait4Shutdown(timeout time.Duration) {
	defaultShutdownCtx.WaitShutdown(timeout)
}

type shutdownContext struct {
	sigChan  chan os.Signal
	status   atomic.Value
	innerCtx context.Context
	cleans   []CleanFunc
}

type CleanFunc func()

func (ctx *shutdownContext) WaitShutdown(timeout time.Duration) {
	select {
	case <-ctx.innerCtx.Done():
		shutdownLogger.Trace("received context done", Dict{"err": ctx.innerCtx.Err().Error()})
	case s := <-ctx.sigChan:
		shutdownLogger.Trace("received shutdown signal", Dict{"sig": s.String()})
	}
	ctx.status.Store(Shutdowning)
	ctx.waitCleans(timeout)
	shutdownLogger.Trace("shutdown finish...", nil)
}

func (ctx *shutdownContext) waitCleans(timeout time.Duration) {
	shutdownLogger.Trace("start cleaning...", nil)
	select {
	case <-ctx.runCleans():
		shutdownLogger.Trace("shutdown clean all", nil)
	case <-time.After(timeout):
		shutdownLogger.Trace("shutdown timeout, just ignore", Dict{"duration": timeout.String()})
	}
	ctx.status.Store(Shutdowned)
}

func (ctx *shutdownContext) Status() ShutdownStatus {
	return ctx.status.Load().(ShutdownStatus)
}

func (ctx *shutdownContext) Clean(cleanHandler ...CleanFunc) {
	ctx.cleans = append(ctx.cleans, cleanHandler...)
}

func (ctx *shutdownContext) runCleans() chan bool {
	finished := make(chan bool)
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(ctx.cleans))
		for _, clean := range ctx.cleans {
			go func(callback CleanFunc) {
				callback()
				wg.Done()
			}(clean)
		}
		wg.Wait()
		finished <- true
	}()
	return finished
}

func WithShutdownContext(ctx context.Context) *shutdownContext {
	shutdownCtx := &shutdownContext{
		sigChan:  make(chan os.Signal, 1),
		status:   atomic.Value{},
		innerCtx: ctx,
		cleans:   make([]CleanFunc, 0),
	}
	shutdownCtx.status.Store(Running)
	return shutdownCtx
}
