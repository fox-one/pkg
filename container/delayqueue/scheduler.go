package delayqueue

import (
	"sync"
	"time"
)

type Scheduler interface {
	Retry(after time.Duration)
	ExecuteCount() int64
}

type scheduler struct {
	dur   time.Duration
	retry bool
	mux   sync.Mutex
	idx   int64
}

func (s *scheduler) Retry(after time.Duration) {
	s.mux.Lock()
	if !s.retry {
		s.dur = after
		s.retry = true
	} else if after < s.dur {
		s.dur = after
	}
	s.mux.Unlock()
}

func (s *scheduler) ExecuteCount() int64 {
	return s.idx
}

func (s *scheduler) next() (next time.Duration, ok bool) {
	return s.dur, s.retry
}
