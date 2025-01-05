package area

import "math"

type Rectangle struct {
	Width  float32
	Height float32
}

func (rectangle Rectangle) Area() float32 {
	return rectangle.Height * rectangle.Width
}

type Circle struct {
	Radius float32
}

func (circle Circle) Area() float32 {
	return math.Pi * circle.Radius * circle.Radius
}

type Shape interface {
	Area() float32
}
