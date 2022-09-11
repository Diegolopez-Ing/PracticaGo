package main

import (
	"fmt"
)

func main() {

	et := []int{41, 42, 43, 44, 45, 46, 47, 48, 49}
	fmt.Println(et)
	fmt.Println(len(et))
	fmt.Println(cap(et))

	et = append(et, 49, 50, 51, 52) //Un nuevo arreglo se asignó
	fmt.Println(et)
	fmt.Println(len(et))
	fmt.Println(cap(et))

	y := append(et[:3], et[5:]...) //Un nuevo arreglo se asignó
	fmt.Println(y)
	fmt.Println(et)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	fmt.Println(et)
	fmt.Println(len(et))
	fmt.Println(cap(et))
}
