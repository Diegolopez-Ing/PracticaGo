package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	// Declaring a variable b of type pointer to an int and assigning it the memory address of a.
	var b *int = &a
	fmt.Println(b)
	// Dereferencing the pointer b.
	fmt.Println(*b)
	*b = 43
	fmt.Println(*b)
	fmt.Println(a)

}
