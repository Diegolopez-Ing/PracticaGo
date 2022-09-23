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

// A method.
func (a agenteSecreto) presesentarse() {
	fmt.Println("Hola, soy ", a.nombre, a.apellido, "EL agente secreto se presenta")
}

// A method.
func (p persona) presesentarse() {
	fmt.Println("Hola soy", p.nombre, p.apellido, "la peronsa se presenta")
}

// A humano is a type that can present itself.
// @property presesentarse - This is the method signature.
type humano interface {
	presesentarse()
}

// Bar takes a humano as an argument
func bar(h humano) {
	fmt.Println("Fui pasado a la función bar", h)
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

	/* 	p := persona{
		nombre:   "CAAL",
		apellido: "Guzman",
	} */
	fmt.Println(as1)
	as1.presesentarse()
	as2.presesentarse()

	bar(as1)
	bar(as2)
	//bar(p)
}
