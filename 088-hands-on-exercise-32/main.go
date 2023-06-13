package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	age1 := m["James"]
	fmt.Println(age1)

	if v, ok := m["James"]; ok {
		fmt.Println("There is a BOND entry, and here is the value of his age is ", v)
	}

	age2 := m["Q"]
	fmt.Println(age2)

	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no Q, and here is the zero value of an int", v)
	}
}
