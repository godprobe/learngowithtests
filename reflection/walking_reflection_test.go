// "
// From Twitter
// golang challenge: write a function
// `walk(x interface{}, fn func(string))`
// which takes a struct `x` and calls `fn`
// for all strings fields found inside.
// difficulty level: recursively.
// " (sic)

package reflection

import "testing"

func TestWalk(t *testing.T) {

	expected := "Roger the gopher"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d, want %d", len(got), 1)
	}
}
