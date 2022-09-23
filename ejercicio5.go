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

	fmt.Println(p1.nombre)
	fmt.Println(p1.apellido)

	for i, v := range p1.saboresFav {
		fmt.Println("\t", i, v)
	}

	for i, v := range p2.saboresFav {
		fmt.Println("\t", i, v)

	}

}
