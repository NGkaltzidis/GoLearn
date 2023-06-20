package main

import (
	"fmt"
	"strings"
)

func remove(slice []int, s int) []int {
	return append(slice[:1], slice[s+1:]...)
}

func main() {
	s := 7
	// Deleting from a slice
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	xs := []string{"My", "name", "is", "test", "Nikos"}
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-------")

	//xi = append(xi[:2], xi[5:]...)
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-------")

	fmt.Println(remove(xi, s))
	xs = append(xs[:3], xs[4:]...)

	// convert slice to string
	fmt.Println(strings.Join(xs, " "))

	// Note : [inclusive:exclusive]
	// Whatever is inclusive remains in the slice
	// Whatever is exclusive is removed from the slice
}
