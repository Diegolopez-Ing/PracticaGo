package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}

	fmt.Println(x)
	// fmt.Println(x[incia:hasta])
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x[:])
	fmt.Println(x[1:3])
	fmt.Println(x[2:4])

	for i, v := range x {
		fmt.Println(i, " ", v)
	}

}
