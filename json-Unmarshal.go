package main

import (
	"encoding/json"
	"fmt"
)

type personaUnm struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellidos"`
	Edad     int    `json:"Edad"`
}

func main() {
	s := `[{"Nombre":"Diego","Apellido":"Lopez","Edad":23},{"Nombre":"Alberto","Apellido":"Murillo","Edad":24}]`
	bs := []byte(s)

	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))

	var personas []personaUnm

	err := json.Unmarshal(bs, &personas)
	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println("Toda la dara", personas)
	for i, v := range personas {
		fmt.Println(i, v)
	}

}
