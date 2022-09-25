package main

import "fmt"

type personal struct {
	Nombre string
}

func main() {
	p1 := personal{
		Nombre: "AerosMith",
	}
	fmt.Println(p1)
	cambiame(&p1)
	fmt.Println(p1)
}

// Cambiame takes a pointer to a personal as an argument, and changes the value that the pointer points
// to.
func cambiame(p *personal) {
	p.Nombre = "Jhon Dalton"
	//(*p).Nombre = "Jhon Dalton"
}
