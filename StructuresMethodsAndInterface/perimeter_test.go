package StructuresMethodsAndInterface

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{Height: 9.0, Width: 5.0}
		want := 45.0

		checkArea(t, rectangle, want)
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{Radius: 10}
		want := 314.1592653589793

		checkArea(t, circle, want)
	})

}
