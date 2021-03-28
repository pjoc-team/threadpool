package threadpool

import (
	"context"
	"errors"
)

var (
	// ErrPoolSizeMustPresent pool size is illegal
	ErrPoolSizeMustPresent = errors.New("pool size is illegal")
	// ErrPoolIsFull pool is full
	ErrPoolIsFull = errors.New("pool is full")
)

// ThreadPool thread pool interface
type ThreadPool interface {
	// Go async invoke, returns error when the pool is full.
	Go(func()) error
	// Run sync invoke, pending the func when pool is full.
	Run(func())
}

type threadPool struct {
	wc  chan func()
	ctx context.Context
}

func (t *threadPool) Go(f func()) error {
	select {
	case t.wc <- f:
		return nil
	default:
		return ErrPoolIsFull
	}
}

func (t *threadPool) Run(f func()) {
	t.wc <- f
}

func (t *threadPool) watch() {
	for {
		select {
		case f := <-t.wc:
			go func() {
				f()
			}()
		case <-t.ctx.Done():
			return
		}
	}
}

// NewPool create thread pool with size
func NewPool(ctx context.Context, size int) (ThreadPool, error) {
	if size <= 0 {
		return nil, ErrPoolSizeMustPresent
	}
	wc := make(chan func(), size)
	p := &threadPool{
		wc:  wc,
		ctx: ctx,
	}
	go p.watch()
	return p, nil
}
