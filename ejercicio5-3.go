package main

import "fmt"

type vehiculo struct {
	puerta int
	color  string
}
type camion struct {
	vehiculo
	cuatroRuedas bool
}
type sedan struct {
	vehiculo
	lujos bool
}

func main() {

	Camion := camion{
		vehiculo: vehiculo{
			puerta: 2,
			color:  "Amarillo",
		},
		cuatroRuedas: false,
	}

	Camion2 := sedan{
		vehiculo: vehiculo{
			puerta: 3,
			color:  "Fuxcisa",
		},
		lujos: true,
	}

	fmt.Println("", Camion)
	fmt.Println("", Camion2)

}
