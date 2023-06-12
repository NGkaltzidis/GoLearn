package main

import "fmt"

func main() {
	// Variety of for loops
	for i := 0; i < 5; i++ {
		fmt.Printf("counting to five: %v \t first for loop\n", i)
	}

	y := 5

	for y < 10 {
		fmt.Printf("y is %v \t\t\t second for loop\n", y)
		y++
	}

	for {
		fmt.Printf("y is %v \t\t third for loop\n", y)
		if y > 20 {
			break
		}
		y++
	}

	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("counting even numbers: ", i)
	}
}
