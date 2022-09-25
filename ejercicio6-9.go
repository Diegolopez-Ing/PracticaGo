package main

import "fmt"

func foo(b func(number []int) int, num []int) int {
	n := b(num)
	n++
	return n + 1
}

func main() {
	g := func(number []int) int {
		x := 0
		if len(number) == 0 {

			x = 0
		}
		if len(number) == 1 {

			x = number[0]
		}

		if len(number) != 0 && len(number) != 1 {
			x = number[0] + number[len(number)-1]
		}
		return x

	}
	x := foo(g, []int{1, 2, 3, 4, 5})
	fmt.Println(x)
}
