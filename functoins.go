package main

import "fmt"

func main() {

	foo()
	bar("Boond")
	s := woo("Money pény")
	fmt.Println(s)

	x, y := saludar("Diego", "Lopez")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {

	fmt.Println("Hola desde Foo")
}

func bar(s string) {

	fmt.Println("Hola", s)
}

func woo(st string) string {
	return fmt.Sprint("Retorno desde Woo, ", st)
}

func saludar(n string, a string) (string, bool) {
	p := fmt.Sprint(n, " ", a, `dice "hola".`)
	q := true
	return p, q

}

//Func (r receptor) identificador (parametros) retorno(s) (código)
