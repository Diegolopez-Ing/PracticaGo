package main

import (
	"encoding/json"
	"fmt"
)

type personaMar struct {
	Nombre   string
	Apellido string
	Edad     int
}

func main() {
	p1 := personaMar{
		Nombre:   "Diego",
		Apellido: "Lopez",
		Edad:     23,
	}

	p2 := personaMar{
		Nombre:   "Alberto",
		Apellido: "Murillo",
		Edad:     24,
	}

	personas := []personaMar{p1, p2}
	fmt.Println(personas)

	bs, err := json.Marshal(personas)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(string(bs))
}
