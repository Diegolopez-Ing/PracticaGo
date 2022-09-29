package main

import "fmt"

func main() {
	//Buffered chanel
	// Creating a buffered channel of type int with a capacity of 1.
	ca := make(chan int)
	ca1 := make(chan int, 2)
	ca2 := make(<-chan int, 2) //Sen Only
	ca3 := make(chan<- int, 2) //Received Only
	ca1 <- 42
	ca1 <- 43
	fmt.Println(<-ca1)
	fmt.Println(<-ca1)
	fmt.Println("==================================")
	fmt.Printf("%T\n", ca)
	fmt.Printf("%T\n", ca1)
	fmt.Printf("%T\n", ca2)
	fmt.Printf("%T\n", ca3)

	//Genberal a específico convierte
	fmt.Printf("%T\n", ca3)
	fmt.Printf("c\t%T\n", (<-chan int(ca)))
	fmt.Printf("c\t%T\n", (chan<- int(ca)))

	//Específico a general no convierte
	fmt.Printf("%T\n", ca3)
	fmt.Printf("c\t%T\n", (chan int(ca2)))
	fmt.Printf("c\t%T\n", (chan int(ca3)))

}
