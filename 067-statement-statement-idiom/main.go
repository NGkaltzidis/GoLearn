package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 40
	y := 5
	fmt.Printf("x = %v \ny = %v\n", x, y)

	// In this case we keep "z" scope only within the if and else statements which is a good practice
	// to keep scope as little as possible
	// First : short declaration of z assigning a value which is 2 times a random Integer from 0-40 (where 40 is x value)
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is Greater than or Equal to x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is Less than x which is %v\n", z, x)
	}

}
