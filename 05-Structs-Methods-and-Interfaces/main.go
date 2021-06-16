package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

const (
	pi = 3.14
)

type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

func CircleArea(circle Circle) float64 {
	return circle.Radius * pi
}

func main() {
	fmt.Println(Perimeter(Rectangle{Width: 10.0, Height: 10.0}))
	fmt.Println(Area(Rectangle{Width: 12.0, Height: 6.0}))
	fmt.Println(CircleArea(Circle{Radius: 10.0}))
}
