package main

import "fmt"

func main() {
	// create slice
	xi := []int{42, 43, 44}
	fmt.Println(xi)
	fmt.Println("------")

	// Append to slice
	xi = append(xi, 45, 46)
	fmt.Println(xi)
}
