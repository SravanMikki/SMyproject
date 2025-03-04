// Assignment 2: Temperature Converter
// Use these: Methods, Receivers, Pointers
// Problem Statement:
// Create a program that converts temperatures between Celsius and Fahrenheit.
// Requirements:
// 1.⁠ ⁠Define a struct Temperature with a Value (float64).
// 2.⁠ ⁠Implement methods:
//   - ToFahrenheit() → Converts from Celsius to Fahrenheit.
//   - ToCelsius() → Converts from Fahrenheit to Celsius.
//
// 3.⁠ ⁠Use pointers to modify the temperature directly.
// Formula:
// •⁠  ⁠°F = (°C × 9/5) + 32
// •⁠  ⁠°C = (°F - 32) × 5/9
package main

import "fmt"

type Temperature struct {
	Value float64
}

func (t *Temperature) ToFahrenheit() {
	t.Value = (t.Value * 9 / 5) + 32
}

func (t *Temperature) ToCelsius() {
	t.Value = (t.Value - 32) * 5 / 9
}

func main() {
	temp := Temperature{Value: 25}
	fmt.Println("Celsius:", temp.Value)

	temp.ToFahrenheit()
	fmt.Println("Converted to Fahrenheit:", temp.Value)

	temp.ToCelsius()
	fmt.Println("Converted back to Celsius:", temp.Value)
}
