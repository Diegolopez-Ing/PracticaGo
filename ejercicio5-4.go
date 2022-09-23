package main

import "fmt"

func main() {
	s := struct {
		nombre     string
		amigos     map[string]int
		bebidasFav []string
	}{
		nombre: "Dirego",
		amigos: map[string]int{
			"Lopez":  23232,
			"SErgio": 456456,
			"Amador": 9899,
		},
		bebidasFav: []string{
			"agua",
			"Naranjada",
			"Cervezas",
		},
	}

	fmt.Println(s.nombre)
	fmt.Println("\t Amigos:")
	for k, v := range s.amigos {
		fmt.Println("\t\t", k, v)
	}
	fmt.Println("\t Bebidas:")
	for i, v := range s.bebidasFav {
		fmt.Println("\t\t", i, v)
	}
}
