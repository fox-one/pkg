package delayqueue

import (
	"context"
	"sync"
	"time"

	"github.com/RussellLuo/timingwheel"
	"golang.org/x/sync/semaphore"
)

type Queue struct {
	tw *timingwheel.TimingWheel // schedule task

	mu sync.Mutex                     // protects m
	m  map[string]*semaphore.Weighted // lazily initialized
}

func New(tick time.Duration, wheelSize int64) *Queue {
	return &Queue{
		tw: timingwheel.NewTimingWheel(tick, wheelSize),
	}
}

func (q *Queue) Do(ctx context.Context, key string, delay time.Duration, fn func(s Scheduler)) {
	q.do(ctx, key, delay, 1, fn)
}

func (q *Queue) do(ctx context.Context, key string, delay time.Duration, count int64, fn func(s Scheduler)) {
	q.mu.Lock()
	if q.m == nil {
		q.m = make(map[string]*semaphore.Weighted)
		q.tw.Start()
	}

	sem, ok := q.m[key]
	if !ok {
		sem = semaphore.NewWeighted(1)
		q.m[key] = sem
	}
	q.mu.Unlock()

	q.tw.AfterFunc(delay, func() {
		if err := sem.Acquire(ctx, 1); err != nil {
			return
		}

		s := &scheduler{
			idx: count,
		}

		fn(s)
		sem.Release(1)

		if next, ok := s.next(); ok {
			q.do(ctx, key, next, count+1, fn)
		}
	})
}

func (q *Queue) Stop() {
	q.tw.Stop()
}
