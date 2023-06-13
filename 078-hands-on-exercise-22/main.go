package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("The value of x is : %v\n", x)
	fmt.Printf("The value of y is : %v\n", y)

	if x < 4 && y < 4 {
		fmt.Println("Both less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("Both greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is greater than or equal to 4 AND x is less than or equal to 6")
	} else if y != 5 {
		fmt.Println("y is not equal to 5")
	} else {
		fmt.Println("none of the previous cases were met")
	}

}
