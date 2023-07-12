package main

import "fmt"

func main() {
	// This is one way to create a map
	am := map[string]int{
		"Todd":   42,
		"Henry":  17,
		"Padget": 15,
	}

	fmt.Println("The age of Henry was ", am["Henry"])

	for i, v := range am {
		fmt.Printf("Key : %v, Value : %v\n", i, v)
	}
	fmt.Println("-------------------")

	// This is another way to create a map
	na := make(map[string]int)
	na["Lucas"] = 25
	na["Nicholas"] = 30
	na["Sam"] = 20
	fmt.Println("The age of Nicholas is ", na["Nicholas"])

	for i, v := range na {
		fmt.Printf("Key : %v, Value : %v\n", i, v)
	}
}
