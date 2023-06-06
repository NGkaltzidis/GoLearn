package main

import (
	"fmt"
	"github.com/NGkaltzidis/puppy"
)

func main() {
	fmt.Println("Testing")
	fmt.Println(combined())
}

func combined() string {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()

	s5 := s1 + "\n" + s2 + "\n" + s3 + "\n" + s4
	return s5
}
