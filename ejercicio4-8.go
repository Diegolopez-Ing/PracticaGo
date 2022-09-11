package main

import "fmt"

func main() {
	x := map[string][]string{
		"Eduar_tua":    {"Computadoras", "Montañas", "Playas"},
		"Carlos_Ramon": {"Leer", "Comprar", "Música"},
	}
	for i, k := range x {
		fmt.Println(i)
		for j := 0; j < len(k); j++ {

			fmt.Println(k[j])
		}
	}

	x["Diego_lopez"] = []string{"Arroz", "Carro", "Muñeco"}
	for i, k := range x {
		fmt.Println(i)
		for j := 0; j < len(k); j++ {

			fmt.Println(k[j])
		}
	}

	delete(x, "Eduar_tua")
	fmt.Println(x)

}
