package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(200)

	for {
		if x > 100 {
			fmt.Printf("x value is %v, which is higher than 100", x)
			break
		}
		fmt.Printf("x value is : %v\n", x)
		x++
	}
}
