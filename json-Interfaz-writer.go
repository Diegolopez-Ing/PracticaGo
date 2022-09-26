package main

import (
	"fmt"
	"io"
	"os"
)

type personaUnm struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellidos"`
	Edad     int    `json:"Edad"`
}

func main() {
	fmt.Println("Hola Diego")
	fmt.Fprintln(os.Stdout, "Hola Diego")
	io.WriteString(os.Stdout, "Hola Diego")
}
