package main

import "fmt"

var a int

type dinero int

var b dinero

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	//a = b   cannot use b (variable of type dinero) as int value in assignment
	// Converting the value of b from dinero to int.
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
