package main

import (
	"fmt"
)

func main() {
	a := func(b string) string {
		return b
	}("Hola Diego")

	fmt.Println(a)
}
