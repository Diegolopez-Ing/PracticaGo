package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	//fmt.Sprintf("",)//Las primeras comillas indican el formato que tendrá el texto
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Print(s)

}
