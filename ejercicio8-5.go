package main

import (
	"fmt"
	"sort"
)

type usuario struct {
	Nombre   string
	Apellido string
	Edad     int
	Dichos   []string
}

type SortByAge []usuario

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

type SortByLastname []usuario

func (a SortByLastname) Len() int           { return len(a) }
func (a SortByLastname) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByLastname) Less(i, j int) bool { return a[i].Apellido < a[j].Apellido }

func main() {
	u1 := usuario{
		Nombre:   "Eduar",
		Apellido: "Tua",
		Edad:     32,
		Dichos: []string{
			"Cachicamo diciéndole a morrocoy conchudo",
			"La mona, aunque se vista de seda, mona se queda",
			"Poco a poco se anda lejos",
		},
	}

	u2 := usuario{
		Nombre:   "Carlos",
		Apellido: "Pérez",
		Edad:     27,
		Dichos: []string{
			"Camarón que se duerme se lo lleva la corriente",
			"A ponerse las alpargatas que lo que viene es joropo",
			"No gastes pólvora en zamuro",
		},
	}

	u3 := usuario{
		Nombre:   "Che",
		Apellido: "López",
		Edad:     54,
		Dichos: []string{
			"Ni lava ni presta la batea",
			"Hijo de gato, caza ratón",
			"Más vale pájaro en mano que cien volando",
		},
	}

	usuarios := []usuario{u1, u2, u3}

	fmt.Println(usuarios)
	fmt.Println("==================================================================")
	sort.Sort(SortByAge(usuarios))
	for i, v := range usuarios {
		fmt.Println("\tPERSONA #", i+1)
		fmt.Println("\t", v.Nombre, v.Apellido, v.Edad)
		for _, dicho := range v.Dichos {
			fmt.Println("\t\t", dicho)
		}
	}

	fmt.Println("==================================================================")
	sort.Sort(SortByLastname(usuarios))
	for i, v := range usuarios {
		fmt.Println("\tPERSONA #", i+1)
		fmt.Println("\t", v.Nombre, v.Apellido, v.Edad)
		for _, dicho := range v.Dichos {
			fmt.Println("\t\t", dicho)
		}
	}

	// Tu código va aquí

}
