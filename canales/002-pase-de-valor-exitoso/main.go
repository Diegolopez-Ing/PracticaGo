package main

import "fmt"

func main() {
	//unbuffered chanel
	// Creating a channel of type int.
	ca := make(chan int)
	ca <- 42
	fmt.Println(<-ca)
}
