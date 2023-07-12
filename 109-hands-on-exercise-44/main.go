package main

import "fmt"

func main() {
	xs := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	xs1 := append(xs[:5])
	xs2 := append(xs[5:])
	xs3 := append(xs[2:7])
	xs4 := append(xs[1:6])

	fmt.Print(xs1, "\n", xs2, "\n", xs3, "\n", xs4, "\n")
}
