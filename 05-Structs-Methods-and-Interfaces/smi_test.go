package main

import "testing"

func TestArea(t *testing.T) {

	t.Run("Rectangles", func(t *testing.T) {
		rect := Rectangle{12.0, 6.0}
		got := Area(rect)
		want := 72.0

		if got != want {
			t.Errorf("got %.2f, but want %.2f", got, want)
		}
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := CircleArea(circle)
		want := 31.400000000000002

		if got != want {
			t.Errorf("got %.2f, but want %.2f", got, want)
		}
	})

}
