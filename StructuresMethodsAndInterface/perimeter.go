package StructuresMethodsAndInterface

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Height * t.Base) / 2
}
