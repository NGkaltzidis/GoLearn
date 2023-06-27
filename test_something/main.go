package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}
	xi = append(xi[1:], xi[2:]...)
	fmt.Println(xi)
}
