package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Mi primera expresio Funcion")
	}
	f()

	g := func(x int) {
		fmt.Println("EL año que se descubrió América fué:", x)
	}

	g(1492)
}
