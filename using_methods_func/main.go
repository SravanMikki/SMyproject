// You’re tasked with building a simple gradebook system in Golang to track student information and calculate their average grades. The system should allow you to:

// Store student details (name and a list of grades).
// Add a new grade to a student’s record.
// Calculate the student’s average grade.
// Update the student’s name if needed.
// Requirements:
// Use a struct to represent a Student with a name (string) and grades (slice of float64).
// Use a function with a pointer to add a grade to the student’s grade list (modifying the original data).
// Use a function to calculate the average grade (returning a float64).
// Use a function with a pointer to update the student’s name.
// In main(), demonstrate all these operations for one student.//
package main

import "fmt"

type Dog struct {
	Name string
}

func (d Dog) bark() {
	fmt.Println(d.Name, "says: bow bow!")
}

func main() {
	d := Dog{Name: "Tommy"}
	d.bark()
}
