package main

import "fmt"

func main() {

	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := foo(b...)
	c := bar(b)
	fmt.Println(a)
	fmt.Println(c)

}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
