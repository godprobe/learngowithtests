package async_counter

import "testing"

// Concurrency-safe counter.

func TestAsyncCounter(t *testing.T) {
	t.Run("incrementing the counter three times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})
}

func assertCount(t *testing.T, c Counter, count int) {
	if c.Value() != count {
		t.Errorf("got %d, want %d", c.Value(), count)
	}
}
