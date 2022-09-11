package main

import (
	"fmt"
)

func main() {

	m := map[string]int{
		"Batman": 32,
		"RObin":  24,
	}
	fmt.Println(m)
	fmt.Println(m["Batman"])
	fmt.Println(m["RObin"])
	fmt.Println(m["Diego"])
	v, ok := m["Batman"]
	fmt.Println(v, ok)

	if v, ok := m["RObin"]; ok {
		fmt.Println("Imprimiendo desde el ok", v, ok)

	}

}
