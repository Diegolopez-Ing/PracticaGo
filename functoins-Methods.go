package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
}

type agenteSecreto struct {
	persona
	lpm bool
}

func (a agenteSecreto) presesentarse() {
	fmt.Println("Hola, soy ", a.nombre, a.apellido)
}

func main() {

	as1 := agenteSecreto{
		persona: persona{
			nombre:   "Diego",
			apellido: "Lopez",
		},
		lpm: true,
	}

	as2 := agenteSecreto{
		persona: persona{
			nombre:   "Alberto",
			apellido: "Murillo",
		},
		lpm: false,
	}
	fmt.Println(as1)
	fmt.Println(as2)
	as1.presesentarse()
	as2.presesentarse()
}
