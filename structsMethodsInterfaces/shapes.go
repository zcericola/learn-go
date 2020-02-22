package main

//Rectangle represents the shape rectangle with Width and Length properties
type Rectangle struct {
	Width  float64
	Height float64
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
