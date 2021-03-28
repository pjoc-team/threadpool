package threadpool

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func Test_threadPool_Go(t1 *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	pool, err := NewPool(ctx, 100)
	if err != nil {
		t1.Fatal(err.Error())
	}
	loopSize := 1000
	for i := 0; i < loopSize; i++ {
		err := pool.Go(
			func() {
				time.Sleep(100 * time.Microsecond)
			},
		)
		if err != nil && err != ErrPoolIsFull {
			t1.Fatal(err.Error())
		}
	}
}

func Test_threadPool_Run(t1 *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	pool, err := NewPool(ctx, 100)
	if err != nil {
		t1.Fatal(err.Error())
	}
	var count int64 = 0
	loopSize := 1000
	wg := &sync.WaitGroup{}
	for i := 0; i < loopSize; i++ {
		wg.Add(1)
		pool.Run(
			func() {
				defer wg.Done()
				time.Sleep(100 * time.Microsecond)
				atomic.AddInt64(&count, 1)
			},
		)
	}
	wg.Wait()
	if int(count) != loopSize {
		t1.FailNow()
	}
}
