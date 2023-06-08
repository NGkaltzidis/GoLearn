package main

import "fmt"

func main() {
	x := 40
	y := 5
	fmt.Printf("x = %v \ny = %v\n", x, y)

	switch {
	case x < 42:
		fmt.Println("x is Less than 42")
	case x == 42:
		fmt.Println("x is Equal to 42")
	case x > 42:
		fmt.Println("x is Greater than 42")
	default:
		fmt.Println("this is the default case for x")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("x is 41")
	case 42:
		fmt.Println("x is 42")
	case 43:
		fmt.Println("x is 43")
	default:
		fmt.Println("This is the default case for x")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of fallthrough: x is 41")
	case 42:
		fmt.Println("printed because of fallthrough: x is 42")
	case 43:
		fmt.Println("printed because of fallthrough: x is 43")
	default:
		fmt.Println("printed because of fallthrough: this is the default case for x")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of fallthrough: x is 41")
		fallthrough
	case 42:
		fmt.Println("printed because of fallthrough: x is 42")
		fallthrough
	case 43:
		fmt.Println("printed because of fallthrough: x is 43")
		fallthrough
	default:
		fmt.Println("printed because of fallthrough: this is the default case for x")
	}
}
