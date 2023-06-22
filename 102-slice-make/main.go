package main

import "fmt"

func main() {
	// slice of strings

	xs := []string{"a", "b", "c", "d"}
	fmt.Println(xs)

	xi := make([]int, 2, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))

	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xi)
	fmt.Println("-----------")

	xi = append(xi, 10, 11, 12, 13)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("-----------")

}
