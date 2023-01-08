package perimeter

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	X float64
	Y float64
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.X + rectangle.Y) * 2
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.X * rectangle.Y
}

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return circle.Radius * circle.Radius * math.Pi
}
