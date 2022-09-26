package main

import (
	"fmt"
	"sort"
)

type Persona struct {
	Nombre string
	Edad   int
}

type AgeBy []Persona

func (a AgeBy) Len() int           { return len(a) }
func (a AgeBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AgeBy) Less(i, j int) bool { return a[i].Nombre < a[j].Nombre }

func main() {

	p1 := Persona{
		Nombre: "Diego",
		Edad:   28,
	}
	p2 := Persona{
		Nombre: "Maria",
		Edad:   25,
	}
	p3 := Persona{
		Nombre: "Carolina",
		Edad:   32,
	}
	p4 := Persona{
		Nombre: "Adrinaa",
		Edad:   60,
	}

	personas := []Persona{p1, p2, p3, p4}
	fmt.Println(personas)
	sort.Sort(AgeBy(personas))

	fmt.Println(personas)
}
