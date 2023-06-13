package main

import "fmt"

func main() {
	// map with string keys and int values
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	// for range loop to print keys and values of a map
	for i, x := range m {
		fmt.Printf("The key %v has value %v\n", i, x)
	}
}
