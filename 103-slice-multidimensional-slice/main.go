package main

import "fmt"

func main() {
	// create two slices
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Monneypenny", "Guiness", "Wolverine Tracks"}

	fmt.Println(jb)
	fmt.Println(jm)

	// create a multidimensional slice
	xxs := [][]string{jb, jm}
	fmt.Println(xxs)
}
