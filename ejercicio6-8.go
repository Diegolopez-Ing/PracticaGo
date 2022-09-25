package main

import (
	"fmt"
)

func myName() func() int {
	return func() int {
		return 42
	}
}

func main() {
	a := myName()

	fmt.Printf("%T\n", a)
	fmt.Println(a())
}
