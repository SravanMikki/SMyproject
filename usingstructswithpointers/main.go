// Imagine we have a user profile system, and we need a function to update a user's age. Instead of passing the user struct by value (which would create a copy), we use pointers to modify the actual user data.

package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func UpdateAge(u *User, age int) {
	u.Age = age

}
func main() {
	User1 := User{Name: "Mikki", Age: 24}
	fmt.Println("Before update:", User1.Name, "Age:", User1.Age)
	UpdateAge(&User1, 23)
	fmt.Println("After update:", User1.Name, "Age:", User1.Age)
}
