package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{3.25, 12.5}
	got := Perimeter(rectangle)
	want := 31.5

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{3.25, 12.5}
	got := Area(rectangle)
	want := 40.625

	if got != want {
		t.Errorf("got %.3f, want %.3f", got, want)
	}
}
