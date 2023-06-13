package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is from init func")

}

func main() {
	x := rand.Intn(400)
	fmt.Printf("The value of x is %v\t", x)

	switch {
	case x <= 100:
		fmt.Println("Between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("Between 101 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("Between 201 and 250")
	default:
		fmt.Println("This was more than 250")
	}
}
