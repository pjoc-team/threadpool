package threadpool

import "errors"

var (
	// ErrPoolSizeMustPresent pool size is illegal
	ErrPoolSizeMustPresent = errors.New("pool size is illegal")
)

// ThreadPool thread pool interface
type ThreadPool interface {
	// Go async invoke, returns error when the pool is full.
	Go(func()) error
	// Run sync invoke, pending the func when pool is full.
	Run(func())
}

type threadPool struct {
	wc chan func()
}

func (t threadPool) Go(f func()) error {
	panic("implement me")
}

func (t threadPool) Run(f func()) {
	panic("implement me")
}

// NewPool create thread pool with size
func NewPool(size int) (ThreadPool, error) {
	if size <= 0 {
		return nil, ErrPoolSizeMustPresent
	}
	wc := make(chan func(), size)
	p := &threadPool{
		wc: wc,
	}
	return p, nil
}
