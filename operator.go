package main
import "fmt"

var z = 41

func main() {
	var w int
	y := "James Bond"
	x := 42 + 7
	fmt.Println(y)
	fmt.Println(x)
	x = 50
	fmt.Println(x)
	fmt.Println(w)
	numero()
}

func numero() {
	fmt.Println(z)
}
