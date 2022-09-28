package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `Clave123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 4)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(s)
	fmt.Println(bs)
	claveLogin := `clave123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(claveLogin))
	if err != nil {
		fmt.Println("No te puede Loguear")
		return
	}
	fmt.Println("Te haz logueado")
}
