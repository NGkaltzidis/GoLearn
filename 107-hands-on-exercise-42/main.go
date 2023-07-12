package main

import "fmt"

func main() {
	// Create an array of 5 indexes and assign values to each index
	va := [5]int{1, 2, 3, 4, 5}

	sa := [5]int{}

	// Create the array as alternative way using a for loop
	for i := 0; i < 5; i++ {
		sa[i] = i
	}

	// Tutor's way
	for i, v := range sa {
		fmt.Printf("Value : %v,  Type : %T, Index : %v\n", v, v, i)
	}

	// Using range loop through the array and print index - value - type
	for i := range va {
		fmt.Printf("Value of index %v is : %v , and type of value is %T\n", i, va[i], va[i])
	}

}
