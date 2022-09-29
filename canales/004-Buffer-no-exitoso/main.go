package main

import "fmt"

func main() {
	//Buffered chanel
	// Creating a buffered channel of type int with a capacity of 1.
	ca := make(chan int, 1)
	ca <- 42
	ca <- 43
	fmt.Println(<-ca)
}
