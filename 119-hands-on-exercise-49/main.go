package main

import "fmt"

func main() {
	// Create a map that has key = string, value = slice of string
	am := map[string][]string{}
	xs := []string{"shaken, not stirred", "martinls", "fast cars"}
	xxs := []string{"intelligence", "literal", "computer science"}
	xi := []string{"cats", "ice-cream", "sunset"}
	am["bond_james"] = xs
	am["moneypenny_jenny"] = xxs
	am["no_dr"] = xi

	// This is an alternative way to create the same map
	as := map[string][]string{
		"bond_james":       {"shaken, not stirred", "martinls", "fast cars"},
		"moneypenny_jenny": {"intelligence", "literal", "computer science"},
		"no_dr":            {"cats", "ice-crea", "sunset"},
	}

	fmt.Println("FIRST : --------------------------")
	// Range in map to print Key - Value pairs and then Index - Value of slice
	for k, v := range am {
		fmt.Printf("Key : %v, Value : %v\n", k, v)
		for i, k := range v {
			fmt.Printf("Index : %v, Value : %v\n", i, k)

		}
		fmt.Println("---------")
	}

	fmt.Println("SECOND : --------------------------")
	// Range on the second map
	for k, v := range as {
		fmt.Printf("Key : %v, Value : %v\n", k, v)
		for i, k := range v {
			fmt.Printf("Index : %v, Value : %v\n", i, k)

		}
		fmt.Println("---------")
	}
}
