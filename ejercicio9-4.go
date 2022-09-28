package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	incremento := 0
	gs := 100
	wg.Add(gs)
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := incremento
			runtime.Gosched()
			v++
			incremento = v
			fmt.Println(incremento)
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("El valor del incremento es: ", incremento)
	fmt.Println(incremento)

}
