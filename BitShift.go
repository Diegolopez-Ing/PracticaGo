package main

import "fmt"

const (
	// Using iota to create a series of constants.
	_  = iota
	kb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
	tb = 1 << (iota * 10)
)

func main() {

	fmt.Printf("%d\t\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t\t%b\n", gb, gb)
	fmt.Printf("%d\t\t\t\t%b\n", tb, tb)

}
