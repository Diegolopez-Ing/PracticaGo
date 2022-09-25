package main

import "fmt"

func main() {

	a := foo()
	fmt.Println(a)

	s, x := bar()

	fmt.Println(s)

	fmt.Println(x)

}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 2, "Hola"
}

//Func (r receptor) identificador (parametros) retorno(s) (código)
