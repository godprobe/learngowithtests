package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{3.25, 12.5}
	got := rectangle.Perimeter()
	want := 31.5

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{3.25, 12.5}
		got := rectangle.Area()
		want := 40.625

		if got != want {
			t.Errorf("got %.3f, want %.3f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{5.625}
		got := circle.Area()
		want := 99.401955054989551685732075799078

		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	})
}
