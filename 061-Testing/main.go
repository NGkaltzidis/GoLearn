package main

import "fmt"

// The init function runs before main, if used. In every other case / scenario, main function is the very first
// component that runs on the program.
func init() {
	fmt.Println("Test1")
}

func main() {
	fmt.Println("Test2")
}
