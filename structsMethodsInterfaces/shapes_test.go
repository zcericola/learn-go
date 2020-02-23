package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want) //%.2f will format a float64 to 2 decimal places
	}
}

func TestArea(t *testing.T) {
	/*areaTests is an anonymous struct which is an array of structs each containing a name, shape, and the hasArea property
	this is an example of table based tests which make it easy to test multiple cases
	, but are cumbersome when only simple testing is needed
	tests should act as an assertion of truth about the code, not as a sequence of operations*/

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		//naming struct fields is optional but helps readability
		{shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}
}
