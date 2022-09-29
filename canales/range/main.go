package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	//Enviar
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Finalizando")
}
