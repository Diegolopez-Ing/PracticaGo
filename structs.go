package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func main() {

	// Creating a new instance of the struct persona and assigning it to the variable pl.
	pl := persona{
		nombre:   "Diego Alberto",
		apellido: "Lopez Murillo",
		edad:     25,
	}
	p2 := persona{
		nombre:   "Alejandro",
		apellido: "Martinez",
		edad:     18,
	}

	fmt.Println(pl)
	fmt.Println(pl.edad)
	fmt.Println(p2.apellido)
	fmt.Printf("%T\n", pl)
}
