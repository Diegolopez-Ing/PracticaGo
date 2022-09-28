package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Sistema Operativo:%v\nArquitectrura:%v\n", runtime.GOOS, runtime.GOARCH)
}
