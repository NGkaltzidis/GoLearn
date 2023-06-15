package main

import "fmt"

func main() {
	xs := []string{"Almond Biscotti Cafe", "Banana Pudding", "Balsamic Strawberry (GF)"}
	fmt.Println(xs)
	fmt.Println("Length of slice : ", len(xs))

	for _, v := range xs {
		fmt.Printf("%v\n", v)
	}

	fmt.Println("--------")
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])
	fmt.Println("--------")

	// Looping to access slice values
	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i])
	}
}
