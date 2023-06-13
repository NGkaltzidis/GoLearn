package main

import (
	"fmt"
	"math/rand"
)

func main() {

	y := 0

	for i := 0; i < 42; i++ {
		x := rand.Intn(5)

		switch x {
		case x:
			fmt.Printf("The value of x is %v\t", x)
		}
		y++
		// counting the iteration times.
		fmt.Printf("Iteration number %v\n\n", y)
	}

}
