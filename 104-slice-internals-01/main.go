package main

import "fmt"

func main() {
	// Since slices are pointing towards the underlined arrays : a points to the array, b points to a
	// so b takes the values from a no matter what
	a := []int{0, 1, 2, 3, 4, 5}
	b := a

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("-------")

	a[0] = 7
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("-------")
}
