package main

import "fmt"

func main() {
	fmt.Println("fsdfsd")
	x1 := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(x1...)
	fmt.Println("El total es ", x)
}

// The function sum takes a slice of ints and returns an int.
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
