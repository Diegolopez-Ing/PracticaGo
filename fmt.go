package main

import "fmt"

var a int
var b string = "program"
var c bool
var d bool = true

func main() {
	e := 42
	f := "Dice hola mundo."
	g := `El doctor dice que comer vegetales es 
	saludable`
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	/*Printf(format string, a ...any) (n int, err error)
	  Printf formats according to a format specifier and writes to standard output.
	 It returns the number of bytes written and any write error encountered. */
	fmt.Printf("EL vaor de la variabla es: %v\n", a)
	fmt.Printf("El vaor de la variabla es: %v\n", b)
	fmt.Printf("El tipo de c es tipo: %T\n", c)
	s1 := fmt.Sprint("El ", b, " Dice Hola mundo")
	fmt.Println(s1)
	fmt.Printf("%T", s1)

}
