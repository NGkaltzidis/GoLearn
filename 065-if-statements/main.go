package main

import "fmt"

func main() {
	x := 40
	y := 5
	fmt.Printf("x = %v \ny = %v\n", x, y)

	if x < 42 {
		fmt.Println("Less than meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than meaning of life")
	} else {
		fmt.Println("Equal to, or greater than, the meaning of life")
	}
	if x < 42 {
		fmt.Println("Less than meaning of life")
	} else if x == 42 {
		fmt.Println("Equal to the meaning of life")
	} else {
		fmt.Println("Greater than the meaning of life")
	}

}
