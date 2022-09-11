package main

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	for j := 0; j <= 10; j++ {
	// 		fmt.Println("VSLOR DE J", j+i)
	// 	}
	// 	fmt.Println("VALOR DE I", i)
	// }

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	//Funciona como el método while
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	y := 1

	for {
		y++
		if y > 100 {
			break
		}
		if y%2 != 0 {
			continue
		}
		fmt.Println(y)
	}

	//Valores Hexadecimal
	for y := 33; y <= 122; y++ {
		fmt.Printf("%#U\t", y)
	}

}
