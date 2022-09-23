package main

import "fmt"

func main() {
	fmt.Println("fsdfs")

	// Deferring the execution of foo() until the surrounding function returns.
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
