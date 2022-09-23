package main

import "fmt"

func main() {
	foo()
	func(x int) {
		fmt.Println("La funcion anonima se ejecutó", x)
	}(42)
}

func foo() {
	fmt.Println("Foo se ejecutó")
}
