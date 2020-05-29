package delayqueue

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

func TestQueue_Do(t *testing.T) {
	ctx := context.Background()
	q := New(50*time.Millisecond, 64)

	wg := sync.WaitGroup{}

	task_1 := func(s Scheduler) {
		t.Log("task 1 start")
		time.Sleep(time.Second)
		t.Log("task 1 end")

		wg.Done()
	}

	task_2 := func(s Scheduler) {
		t.Log("task 2 start")
		time.Sleep(time.Second)
		t.Log("task 2 end")
		wg.Done()
	}

	task_3 := func(s Scheduler) {
		t.Log("task 3 start")
		time.Sleep(time.Second)
		t.Log("task 3 end")

		q.Do(ctx, "", 5*time.Second, task_1)

		wg.Done()
	}

	t.Run("same key", func(t *testing.T) {
		wg.Add(2)
		q.Do(ctx, "", 0, task_1)
		q.Do(ctx, "", 0, task_2)
		wg.Wait()
	})

	t.Run("different key", func(t *testing.T) {
		wg.Add(2)
		q.Do(ctx, "1", time.Second, task_1)
		q.Do(ctx, "2", time.Second, task_2)
		wg.Wait()
	})

	t.Run("retry", func(t *testing.T) {
		wg.Add(4)
		q.Do(ctx, "", 1*time.Second, task_1)
		q.Do(ctx, "", 1*time.Second, task_2)
		q.Do(ctx, "", 1*time.Second, task_3)
		wg.Wait()
	})
}

func TestSchedule(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	q := New(50*time.Millisecond, 64)
	defer q.Stop()

	var count int64

	q.Do(ctx, "", 0, func(s Scheduler) {
		count = s.ExecuteCount()
		t.Log("execute", count)

		if count >= 10 {
			cancel()
			return
		}

		s.Retry(500 * time.Millisecond)
	})

	select {
	case <-ctx.Done():
		time.Sleep(2 * time.Second)
	}

	assert.Equal(t, int64(10), count)
}
