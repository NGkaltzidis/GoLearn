package main

import "fmt"

func main() {
	x := 40
	y := 5
	fmt.Printf("x = %v \ny = %v\n", x, y)

	if x < 42 && y < 42 {
		// requires both conditions to result true in order to run
		fmt.Println("both or less than the meaning of life")
	}

	if x > 30 || y < 42 {
		// requires one of the two conditions to result true in order to run
		fmt.Println("Less than meaning of life")
	}

	if x != 42 && y != 10 {
		fmt.Println("x is not 42 AND y is not 10")
	}
}
