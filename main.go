package main

import (
	"Sgofiles/operations"
	"fmt"
)

func main() {
	a, b := 15, 10
	fmt.Printf("Add: %d\nSub: %.2f\nMultiply: %d\nDivide: %.2f\n",
		operations.Add(a, b), operations.Sub(a, b), operations.Multiply(a, b), operations.Divide(a, b))

}
