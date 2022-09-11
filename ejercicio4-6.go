package main

import "fmt"

func main() {
	x := make([]string, 3, 3)
	x = []string{"Medellin", "Colombia", "Apartadó"}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Printf("%d\t %s\n", i, x[i])
	}
}
