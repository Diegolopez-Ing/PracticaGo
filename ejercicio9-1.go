package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}

func a() {
	for i := 0; i < 20; i++ {
		fmt.Printf("A:%v\t", i)
	}
	wg.Done()

}

func b() {
	for i := 0; i < 20; i++ {
		fmt.Printf("B:%v\t", i)
	}
	fmt.Println("================================")
	wg.Done()
}
