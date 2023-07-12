package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// 42, 43, 44, 48, 49, 50, 51

	// starting slicing the slice. 1. first portion from 42-44, then second portion from 48 onwards
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
