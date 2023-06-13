package main

import "fmt"

func main() {
	// slice xi
	xi := []int{42, 43, 44, 45, 46, 47}

	// has index "i" and value "x"
	for i, x := range xi {
		fmt.Printf("index %v, has value : %v\n", i, x)
	}
}
