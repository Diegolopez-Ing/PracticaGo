package main

import "fmt"

func main() {

	m := map[string]int{
		"Batman": 32,
		"Robin":  24,
	}

	delete(m, "dIEGO")
	fmt.Println(m)

	if v, ok := m["Robin"]; ok {
		fmt.Println("Eliminando la llave con el valor", v)
		delete(m, "Diego")
	}
	fmt.Println(m)
}
