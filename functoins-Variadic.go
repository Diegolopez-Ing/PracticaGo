package main

import "fmt"

func main() {
	fmt.Println("fsdfsd")
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("El total es ", x)
}

// Foo takes a variadic parameter of type int and doesn't return anything.
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Println("%T\n", x)

	suma := 0
	for i, v := range x {
		suma += v
		fmt.Println("El valor en el indice", i, "Suma", v, "Al total, quedano", suma)
	}
	return suma
}
