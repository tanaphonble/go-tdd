package main

import (
	"math"
)

// Shape - A shape interface
type Shape interface {
	Area() float64
}

// Rectangle - A rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Area - Calculate area of rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle - A circle struct
type Circle struct {
	Radius float64
}

// Area - Calculate area of circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter - Calculate perimeter of rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
