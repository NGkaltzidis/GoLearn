package main

import "fmt"

func main() {
	// create slice
	xi := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(xi)
	fmt.Println(len(xi))

	// Slice a slice
	// [inclusive:exclusive]
	fmt.Printf("%#v\n", xi[0:4])
	fmt.Println("---------")
	fmt.Println("---------")

	// [:exclusive]
	fmt.Printf("%#v\n", xi[:3])
	fmt.Println("---------")

	// [inclusive:]
	fmt.Printf("%#v\n", xi[4:])
	fmt.Println("---------")

}
