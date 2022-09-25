package main

import "fmt"

type personaejer struct {
	Apellido string
	Edad     string
	Nombre   string
}

func (per personaejer) presentar() {
	fmt.Println("Hola Soy", per.Nombre, " ", per.Apellido)
}

func main() {

	p := personaejer{
		Nombre:   "Diego",
		Apellido: "Lopez",
		Edad:     "29",
	}

	p.presentar()

}
