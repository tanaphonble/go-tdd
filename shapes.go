package main

// Rectangle - A rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter - Calculate perimeter of rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area - Calcyulate area of rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
