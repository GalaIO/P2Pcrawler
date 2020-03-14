package misc

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

var workPoolLogger = GetLogger().SetPrefix("workpool")

type AsyncWorkFunc func()

type WorkPool struct {
	name       string
	size       int
	queue      chan AsyncWorkFunc
	wg         sync.WaitGroup
	status     atomic.Value
	shutdowned chan bool
	ctx        context.Context
}

func (pool *WorkPool) init() {
	pool.wg.Add(pool.size)
	pool.status.Store(Running)
	go pool.listenCancelLoop()
	for i := 0; i < pool.size; i++ {
		go pool.workerLoop()
	}
	workPoolLogger.Trace("workpool init finished", Dict{"name": pool.name})
}

func (pool *WorkPool) workerLoop() {
	for {
		select {
		case <-pool.shutdowned:
			break
		default:
		}
		worker := <-pool.queue
		worker()
	}
	// not running exit
	pool.wg.Done()
}

func (pool *WorkPool) AsyncSubmit(worker AsyncWorkFunc) {
	if pool.Status() != Running {
		panic(pool.name + " is shutdowned")
	}
	pool.queue <- worker
}

func (pool *WorkPool) Shutdown(timeout time.Duration) {
	if Running != pool.Status() {
		return
	}
	workPoolLogger.Trace("strating shutdown", Dict{"name": pool.name, "timeout": timeout.String()})
	pool.status.Store(Shutdowning)
	// broadcast shutdown
	close(pool.shutdowned)
	select {
	case <-time.After(timeout):
		workPoolLogger.Trace("shutdown timeout just exit", Dict{"name": pool.name})
	case <-pool.waitFinish():
		workPoolLogger.Trace("shutdown success", Dict{"name": pool.name})
	}
	pool.status.Store(Shutdowned)
}

func (pool *WorkPool) waitFinish() chan bool {
	finish := make(chan bool)
	go func() {
		pool.wg.Wait()
		close(finish)
	}()
	return finish
}

func (pool *WorkPool) Status() ShutdownStatus {
	return pool.status.Load().(ShutdownStatus)
}

func (pool *WorkPool) listenCancelLoop() {
	select {
	case <-pool.ctx.Done():
		workPoolLogger.Trace("shutdown from cancel", Dict{"name": pool.name})
		pool.Shutdown(2 * time.Second)
	case <-pool.shutdowned:
		workPoolLogger.Trace("listenCancelLoop exit because shutdown", Dict{"name": pool.name})
		return
	}
}

func NewWorkPool(ctx context.Context, name string, size int) *WorkPool {
	pool := WorkPool{
		name:       name,
		size:       size,
		queue:      make(chan AsyncWorkFunc, size),
		wg:         sync.WaitGroup{},
		status:     atomic.Value{},
		shutdowned: make(chan bool),
		ctx:        ctx,
	}
	pool.init()
	return &pool
}
