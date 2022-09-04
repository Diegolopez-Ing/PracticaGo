package main

import "fmt"

type typetest int

var x typetest
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("EL tipo de y es: %T\n", y)
}
