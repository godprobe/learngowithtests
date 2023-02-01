package async_counter

import (
	"sync"
	"testing"
)

// Concurrency-safe counter.

func TestAsyncCounter(t *testing.T) {
	t.Run("incrementing the counter three times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCount(t, counter, wantedCount)
	})
}

func assertCount(t *testing.T, c *Counter, count int) {
	if c.Value() != count {
		t.Errorf("got %d, want %d", c.Value(), count)
	}
}
