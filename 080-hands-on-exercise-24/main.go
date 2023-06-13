package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("x is : %v \ny is : %v\n", x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("Both less than 4")
		case x > 6 && y > 6:
			fmt.Println("Both greater than 6")
		case x >= 4 && x <= 6:
			fmt.Println("x is greater than or equal to 4 AND x is less than or equal to 6")
		case y != 5:
			fmt.Println("y is not equal to 5")
		default:
			fmt.Println("none of the previous cases were met")
		}
	}
}
