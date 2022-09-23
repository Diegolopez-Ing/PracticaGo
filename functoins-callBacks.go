package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("")

	s := sum(ii...)
	fmt.Println(s)

	s1 := sumaPares(sum, ii...)
	fmt.Println("El total de pares es", s1)
	s2 := sumaImpares(sum, ii...)
	fmt.Println("El total de pares es", s2)

}

func sum(x1 ...int) int {
	fmt.Printf("%T\n", x1)
	total := 0
	for _, v := range x1 {
		total += v
	}
	return total
}

func sumaPares(f func(x1 ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}

// SumaImpares takes a function f that takes a variable number of ints and returns an int, and a
// variable number of ints, and returns an int.
func sumaImpares(f func(x1 ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 != 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}
