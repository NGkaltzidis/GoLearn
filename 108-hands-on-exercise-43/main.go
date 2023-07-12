package main

import "fmt"

func main() {
	xs := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// range in slice to get values and print - VALUES & TYPE
	for i, v := range xs {
		fmt.Printf("Value of index %v : %v , and type : %T\n", i, v, v)
	}

	// %T = Type of specified variable
	// %#v = the whole slice - (type + value)
	// %v = values within the slice
	fmt.Printf("%T \t %#v \t %v \n", xs, xs, xs)

}
