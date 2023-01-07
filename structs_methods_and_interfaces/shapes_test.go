package main

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(3.25, 12.5)
	want := 31.5

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
