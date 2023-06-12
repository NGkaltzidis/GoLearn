package main

import "fmt"

func main() {
	// Modulus
	x := 83 / 40 // this is a division and the result is 2 since 40 x 2 is 80, we have a remained of 3
	y := 83 % 40 // this is a modulus operation usually used to see what is the remainder of the division

	fmt.Println(x, y)

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%v is an even number\n", i)
		}
	}
}
