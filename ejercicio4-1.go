package main

import "fmt"

func main() {
	x := [5]int{}
	fmt.Println(x)
	for i := 0; i < len(x); i++ {
		x[i] = i + 1
	}
	fmt.Println(x)
	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", x)
}
