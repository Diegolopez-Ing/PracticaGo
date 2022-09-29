package main

import "fmt"

func main() {
	//Buffered chanel
	// Creating a buffered channel of type int with a capacity of 2  receive-only.
	ca := make(<-chan int, 2)
	ca <- 42
	ca <- 43
	fmt.Println(<-ca)
	fmt.Println(<-ca)
	fmt.Println("==================================")
	fmt.Printf("%T\n", ca)
}
