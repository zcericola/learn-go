package main

import "math"

//Shape allows to calc the area for any shape
type Shape interface {
	Area() float64
}

//Rectangle represents the shape rectangle with Width and Length properties
type Rectangle struct {
	Width  float64
	Height float64
}

//Area is a method on the Rectangle struct
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

//Circle represents the shape circle with a radius property
type Circle struct {
	Radius float64
}

//Area is a method on the Circle struct
func (c Circle) Area() float64 {
	rSquared := c.Radius * c.Radius
	return math.Pi * rSquared

}

//Triangle represents a triangle with base and height
type Triangle struct {
	Base   float64
	Height float64
}

//Area is method on the Triangle struct
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * .5
}

//Perimeter returns the perimeter length of a shape
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

//Area returns the area of a shape
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

func main() {}
