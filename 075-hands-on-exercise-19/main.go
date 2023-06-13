package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(400)
	fmt.Printf("The value of x is %v\t", x)

	if x <= 100 {
		fmt.Println("Between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("Between 101 and 200")
	} else if x >= 201 && x <= 250 {
		fmt.Println("Between 201 and 250")
	} else {
		fmt.Println("This was more than 250")
	}

	for i := 0; i < 10; i++ {
		y := rand.Intn(3)
		fmt.Println(y)
	}
}
