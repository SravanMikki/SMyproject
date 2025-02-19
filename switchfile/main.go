package main

import "fmt"

func main() {
	age := 30
	switch {
	case age < 18:
		fmt.Println("You're a Minor")
	case age >= 18 && age <= 60:
		fmt.Println("You're an Adult")
	default:
		fmt.Print("You're a Citizen")

	}

}
