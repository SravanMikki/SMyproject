// Assignment 1: Employee Management System
// Use these: Structs, Methods, Pointers, Slices, Maps
// Problem Statement:
// Create a simple Employee Management System that allows adding employees, updating their salary, and displaying their details.
// Requirements:
// 1.⁠ ⁠Define a struct Employee with fields:
//     * ID (int)
//     * Name (string)
//     * Salary (float64)
// 2.⁠ ⁠Implement methods:
//     * UpdateSalary(newSalary float64) → Updates the salary of an employee.
//     * Display() → Prints employee details.
// 3.⁠ ⁠Use a map to store employees (map[int]*Employee).
// 4.⁠ ⁠Implement functions:
//     * AddEmployee(emp *Employee, db map[int]*Employee) → Adds an employee to the map.
//     * RemoveEmployee(id int, db map[int]*Employee) → Removes an employee by ID.

package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func (e *Employee) UpdateSalary(newSalary float64) {
	e.Salary = newSalary
}

func (e Employee) Display() {
	fmt.Printf("ID: %d, Name: %s, Salary: %.2f\n", e.ID, e.Name, e.Salary)
}

func AddEmployee(emp *Employee, db map[int]*Employee) {
	db[emp.ID] = emp
}

func RemoveEmployee(id int, db map[int]*Employee) {
	delete(db, id)
}

func main() {
	employees := make(map[int]*Employee)
	emp1 := &Employee{ID: 1, Name: "Mikki", Salary: 50000}
	emp2 := &Employee{ID: 2, Name: "Alice", Salary: 60000}
	AddEmployee(emp1, employees)
	AddEmployee(emp2, employees)

	emp1.Display()
	emp2.Display()

	emp1.UpdateSalary(55000)

	RemoveEmployee(2, employees)

	emp1.Display()
}
