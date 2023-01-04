package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 7)
	want := "aaaaaaa"
	if got != want {
		t.Errorf("Expected %q, got %q.", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 99)
	}
}

func ExampleRepeat() {
	result := Repeat("z", 6)
	fmt.Println(result)
	// Output: zzzzzz
}
