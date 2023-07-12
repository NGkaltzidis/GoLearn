package main

import (
	"fmt"
)

func main() {
	// create two slices of type string
	// then assign those slices into one multidimensional slice
	xs := []string{"James", "Bond", "Shaken, not stirred"}
	xy := []string{"Miss", "MoneyPenny", "I'm 008 "}
	xxs := [][]string{xs, xy}

	fmt.Println(xxs)

	for i, v := range xxs {
		fmt.Printf("Multidimensional --> Index : %v, Value : %v\n", i, v)
		for a, b := range v {
			fmt.Printf("Inner --> Index : %v, Value : %v\n", a, b)
		}
	}

}
