package main

import "fmt"

// The init function runs before main, if used
func init() {
	fmt.Println("Test1")
}

func main() {
	fmt.Println("Test2")
}
