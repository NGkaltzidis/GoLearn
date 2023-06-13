package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("First loop occuring %v\n\n", i)
		for y := 1; y <= 5; y++ {
			fmt.Printf("Print number %v\n", y)
		}
		fmt.Println("--")
	}
}
