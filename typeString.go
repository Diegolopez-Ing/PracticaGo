package main

import "fmt"

func main() {
	s1 := "Hola soy un string literal interpretado"
	s2 := `Hola soy un string literal no 
						interptreatdo`
	fmt.Println(s1)
	fmt.Printf("EL tipo de s1 es : %T\n", s1)
	fmt.Println(s2)
	fmt.Printf("EL tipo de s2 es : %T\n", s2)
	fmt.Println("")

	bs := []byte(s1)
	fmt.Println(bs)
	// fmt.Printf("%U", bs)

	for i := 0; i < len(s1); i++ {
		fmt.Printf("%#U ", s1[i])
	}
	fmt.Println("")
	for i, v := range s1 {
		fmt.Printf("EL indice de %d el valor es %v\n", i, v)

	}

	fmt.Println("")
	for i, v := range s1 {
		fmt.Printf("EL indice de %d el valor es %q\n", i, v)

	}
}
