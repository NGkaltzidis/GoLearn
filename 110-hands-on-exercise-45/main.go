package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := []int{56, 57, 58, 59, 60}
	x = append(x, 52)

	fmt.Println("52 added \t\t\t", x)

	x = append(x, 53, 54, 55)

	fmt.Println("53, 54, 55 added \t\t", x)

	x = append(x, y...)

	fmt.Println("Slice y added to x\t\t", x)

}
