package main

import "fmt"

type persona struct {
	nombre     string
	apellido   string
	saboresFav []string
}

func main() {

	p1 := persona{
		nombre:     "Diego",
		apellido:   "Lopez",
		saboresFav: []string{"Chantillí", "Vainilla", "Milkey Way"}, //ülitmo elemento lleva ,
	}

	p2 := persona{
		nombre:     "Sergio",
		apellido:   "Algarin",
		saboresFav: []string{"Brownie", "Ron con pasas", "Oreo"},
	}

	m := map[string]persona{
		p1.apellido: p1,
		p2.apellido: p2,
	}

	for k, v := range m {
		fmt.Println("\t", k, v)
		for i, v := range v.saboresFav {
			fmt.Println("\t", i, v)
		}
	}
	fmt.Println("=============================")

	for _, v := range m {
		fmt.Println("\t", v)
	}

}
