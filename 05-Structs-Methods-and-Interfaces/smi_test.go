package main

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 31.41592653589793},
	}

	for _, test := range areaTests {
		got := test.shape.Area()
		if got != test.want {
			t.Errorf("got %g, but want %g", got, test.want)
		}
	}
}

// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g, but want %g", got, want)
// 		}
// 	}

// 	t.Run("Rectangles", func(t *testing.T) {
// 		rect := Rectangle{12.0, 6.0}
// 		checkArea(t, rect, 72.0)
// 	})

// 	t.Run("Circles", func(t *testing.T) {
// 		circle := Circle{10.0}
// 		checkArea(t, circle, 31.41592653589793)
// 	})

// }
