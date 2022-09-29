package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("error de context:\t", ctx.Err())
	fmt.Printf("tipo de context:\t%T\n", ctx)
	fmt.Println("----------")
}
