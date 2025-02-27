package main

import "fmt"

func main() {
	foods := map[string]int{
		"pizza": 10,
		"Cake":  8,
	}
	foods["icecream"] = 10
	delete(foods, "Cake")
	fmt.Println("\nAfter deleting 'Cake':")
	for key, value := range foods {
		fmt.Println(key, ":", value)

	}
}
