package generics

import (
	"testing"
)

// Tests
func TestStack(t *testing.T) {
	t.Run("Integer stack", func(t *testing.T) {
		aStackOfInts := new(StackOfInts)

		// check stack is empty
		AssertTrue(t, aStackOfInts.IsEmpty())

		// add an item, check stack is not empty
		aStackOfInts.Push(451)
		AssertFalse(t, aStackOfInts.IsEmpty())

		// add another item, then remove it
		aStackOfInts.Push(1138)
		value, _ := aStackOfInts.Pop()
		AssertEqual(t, value, 1138)
	})
}

// Helpers
func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
