// Define an interface Shape with one method:
// Area() float64: Returns the area of the shape.
// Define two structs:
// Circle with a field Radius (float64).
// Rectangle with fields Width and Height (float64).
// Implement the Area() method for:
// Circle: Area = π * Radius² (use 3.14 for π).
// Rectangle: Area = Width * Height.
// Write a function:
// CalculateArea(s Shape): Takes a Shape interface and prints its area.
// In main():
// Create a Circle with radius 5.
// Create a Rectangle with width 4 and height 6.
// Call CalculateArea for both shapes.
package main

import "fmt"

// Define the Shape interface
type Shape interface {
	Area() float64
}

// Define Circle struct
type Circle struct {
	Radius float64
}

// Define Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Define Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Define Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Function to calculate and print area
func CalculateArea(s Shape) {
	fmt.Printf("%s Area: %.2f\n", s, s.Area())
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}
	CalculateArea(circle)
	CalculateArea(rectangle)
}
