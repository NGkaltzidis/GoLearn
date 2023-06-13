package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i <= 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("Iteration %v , x is : 3", i)
		}
	}
}
