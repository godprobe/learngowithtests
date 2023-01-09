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

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{3.25, 12.5}, 40.625},
		{Circle{5.625}, 99.401955054989551685732075799078},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g, want %g", got, tt.want)
		}
	}
}
