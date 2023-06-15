package main

import "fmt"

func main() {
	// declare a variable of type [7]int
	var ni [7]int

	// assign a value to index 0
	ni[0] = 42
	fmt.Printf("%#v \t\t\t\t %T\n", ni, ni)

	// array literal
	ni2 := [4]int{55, 56, 57, 58}
	fmt.Printf("%#v \t\t\t\t\t %T\n", ni2, ni2)

	// another way of creating an array literal (In this case compiler counts the elements into the array)
	ni3 := [...]string{"Nikos", "Gkaltzidis"}
	fmt.Printf("%#v \t\t\t %T\n", ni3, ni3)

	fmt.Printf("The length of array ni : %v\n", len(ni))
	fmt.Printf("The length of array ni2 : %v\n", len(ni2))
	fmt.Printf("The length of array ni3 : %v\n", len(ni3))

}
