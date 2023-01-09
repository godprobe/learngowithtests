package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{3.25, 12.5}
	got := rectangle.Perimeter()
	hasPerimeter := 31.5

	if got != hasPerimeter {
		t.Errorf("got %.2f, hasPerimeter %.2f", got, hasPerimeter)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 3.25, Height: 12.5}, hasArea: 40.625},
		{name: "Circle", shape: Circle{Radius: 5.625}, hasArea: 99.401955054989551685732075799078},
		{name: "Triangle", shape: Triangle{Base: 3.25, Height: 12.5}, hasArea: 20.3125},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g, hasArea %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
