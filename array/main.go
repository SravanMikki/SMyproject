package main

import "fmt"

func main() {
	nums := [5]int{-1, 10, 38, 49, -5}
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	fmt.Println(" Minimun number is", min)

}
