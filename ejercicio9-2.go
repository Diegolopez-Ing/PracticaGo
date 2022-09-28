package main

import "fmt"

type PersonaEjer9 struct {
	nombre   string
	apellido string
	Edad     int
}

func (pers *PersonaEjer9) hablar() {
	fmt.Println("Hola : ", pers.nombre)
}

type humano interface {
	hablar()
}

func diAlgo(hu humano) {
	hu.hablar()
}

func main() {
	pers := PersonaEjer9{
		nombre:   "Diego",
		apellido: "LOpez",
		Edad:     21,
	}
	diAlgo(&pers)
	pers.hablar()
	fmt.Println("fdsdf")
}
