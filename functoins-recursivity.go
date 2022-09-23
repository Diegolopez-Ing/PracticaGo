package main

import "fmt"

func main() {

	fmt.Println(4 * 3 * 2 * 1)
	n1 := factorial(4)
	fmt.Println(n1)
	n12 := cicloFact(4)
	fmt.Println(n12)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)

}

func cicloFact(n int) int {
	total := 1

	for ; n > 0; n-- {
		total *= n
	}

	return total
}
