package main

import "fmt"

func main() {
	x := []int{41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	y := x[1:6]
	fmt.Println(y)
	z := x[6:]
	fmt.Println(z)
	w := x[3:8]
	fmt.Println(w)
	v := x[2:7]
	fmt.Println(v)

}
