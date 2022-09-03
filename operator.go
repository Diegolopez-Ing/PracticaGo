package main

import "fmt"

var z = 41

// The main function is the entry point of the program and is executed first.
func main() {
	var w int
	y := "James Bond"
	x := 42 + 7
	fmt.Println(y)
	fmt.Println(x)
	x = 50
	fmt.Println(x)
	fmt.Println(w)
	numero()
}

// Numero() is a function that prints the value of z.
func numero() {
	fmt.Println(z)
}
