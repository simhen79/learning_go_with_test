package structs

import "math"

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

type Shape interface {
	Area() float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.height + r.width)
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) * 0.5
}
