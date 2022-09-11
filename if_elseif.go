package main

import "fmt"

func main() {

	if true {
		fmt.Println("1")
	}
	if false {
		fmt.Println("2")
	}
	if !false {
		fmt.Println("3")
	}
	if !true {
		fmt.Println("4")
	}

	if x := 42; x == 43 {
		fmt.Println("001")
	} else if y := 12; y < 23 {
		fmt.Printf("002")
	}

}
