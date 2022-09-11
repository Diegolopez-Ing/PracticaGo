package main

import "fmt"

func main() {
	x := []string{"Batman", "Jefe", "Vestido de Negro"}
	y := []string{"Robin", "Ayudante", "Vestido de colores"}
	z := [][]string{x, y}

	for v, k := range z {
		fmt.Println(v, k)
		for w, y := range z[v] {
			fmt.Println(w, y)
		}
	}

}
