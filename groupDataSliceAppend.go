package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}
	y := []int{6, 7, 8, 9}

	fmt.Println(x)
	x = append(x, 6, 7)
	fmt.Println(x)
	x = append(x, y...)
	fmt.Println(x)
}
