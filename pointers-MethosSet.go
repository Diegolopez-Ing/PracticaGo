package main

import (
	"fmt"
	"math"
)

type cuadrado struct {
	longitud float64
}

type circulo struct {
	radio float64
}

func (cua cuadrado) area() float64 {
	return cua.longitud * cua.longitud
}

func (cir *circulo) area() float64 {
	return math.Pi * cir.radio * cir.radio
}

type forma interface {
	area() float64
}

func info(frm forma) {
	fmt.Println(frm.area())
}

func main() {

	cir := circulo{
		radio: 12.345,
	}

	cua := cuadrado{
		longitud: 15,
	}

	info(&cir)
	info(cua)
}
