package main

import "fmt"

func main() {
	// This is one way to create a map
	am := map[string]int{
		"Todd":   42,
		"Henry":  17,
		"Padget": 15,
	}

	// Adding key-value to map
	am["Nick"] = 28

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

	// FOR RANGE ON MAP

	// for range over a map ( Example : for key, value := range na {} )
	for k, v := range na {
		fmt.Printf("Key : %v, Value : %v\n", k, v)
	}
	// for range over a map only with value (use underscore for key)
	for _, v := range na {
		fmt.Printf("Value : %v\n", v)
	}
	// for range over a map only with key (use only one variable)
	for k := range na {
		fmt.Printf("Key : %v\n", k)
	}

}
