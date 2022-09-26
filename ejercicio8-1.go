package main

import (
	"encoding/json"
	"fmt"
)

type usuario struct {
	Nombre string
	Edad   int
}

func main() {

	u1 := usuario{
		Nombre: "Diego",
		Edad:   34,
	}
	u2 := usuario{
		Nombre: "Alberto",
		Edad:   23,
	}

	u3 := usuario{
		Nombre: "Elena",
		Edad:   45,
	}

	usuarios := []usuario{u1, u2, u3}
	fmt.Println(usuarios)

	bs, err := json.Marshal(usuarios)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(string(bs))
}
