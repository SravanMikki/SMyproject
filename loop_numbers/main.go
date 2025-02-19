// Loops through numbers from 1 to 10 using a for loop.
// Checks if each number is even or odd using an if-else statement.
// Stops the loop when it reaches the number 9 using a break statement.
// Uses defer to print a message after the loop exits, indicating that the program has finished execution.
package main

import "fmt"

func main() {
	defer fmt.Println("LoopNotdone")

	for i := 0; i <= 15; i++ {
		if i%2 == 0 {
			fmt.Printf("Number %d is Even\n", i)
		} else {
			fmt.Printf("Number %d is Odd\n", i)
		}
		if i == 9 {
			fmt.Println("Breaking out of loop at 9")
			break

		}

	}
}
