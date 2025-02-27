package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

func main() {
	employee := Employee{ID: 001, Name: "Mikki", Age: 24, Department: "Science"}
	fmt.Println("ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Age:", employee.Age)
	fmt.Println("Department:", employee.Department)
}
