package main

import "fmt"

func main() {

	par := make(chan int)
	impar := make(chan int)
	salir := make(chan bool)
	//Enviar
	go enviar(par, impar, salir)

	//Recibir
	recibir(par, impar, salir)

	fmt.Println("Finalizando")
}

func enviar(p, i chan<- int, s chan<- bool) {

	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			i <- j
		}
	}
	close(s)

}

func recibir(p, i <-chan int, s <-chan bool) {
	for {
		select {
		case v := <-p:
			fmt.Println("Desde el canal par", v)
		case v := <-i:
			fmt.Println("Desde el canal impar", v)
		case i, ok := <-s:
			if !ok {
				fmt.Println("Desde com ok", i, ok)
				return
			} else {
				fmt.Println("Desde com ok", i)

			}
		}
	}
}

//COMA OK

/* package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok) ///42 true
} */
