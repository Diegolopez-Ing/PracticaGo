package main

import "fmt"

func main() {
	x := 18
	switch x {
	case 12:
		fmt.Println("12")

	case 23, 18, 24:
		fmt.Println("23")
	case 8:
		fmt.Println("8")
	default:
		fmt.Println("Falseo")
		break

	}
}
