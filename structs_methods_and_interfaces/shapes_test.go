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

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{3.25, 12.5}
		checkArea(t, rectangle, 40.625)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{5.625}
		checkArea(t, circle, 99.401955054989551685732075799078)
	})
}
