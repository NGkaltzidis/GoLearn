package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  17,
		"Padget": 15,
	}

	// Adding key-value to map
	am["Nick"] = 28
	am["Eric"] = 30
	am["Batto"] = 40

	// To delete an element from a map - a Build-In function exists called "delete"
	//delete(am, "Nick")

	fmt.Println("The age of Henry was ", am["Henry"])

	for i, v := range am {
		fmt.Printf("Key : %v, Value : %v\n", i, v)
	}

	fmt.Println("------------------")

	/*v, ok := am["Nick"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key does not exist")
	}*/

	// The best way to use comma ok idiom is the following
	if v, ok := am["Nick"]; !ok {
		fmt.Println("Key does not exist")
	} else {
		fmt.Println("Key exists with a value of :", v)
	}

}
