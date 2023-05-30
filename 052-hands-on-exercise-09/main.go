package main

import "fmt"

var name = "Nikos"

const surname = "Gkaltzidis"

func main() {
	combined := name + " " + surname
	fmt.Println(combined)
}
