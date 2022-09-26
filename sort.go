package main

import (
	"fmt"
	"sort"
)

func main() {
	x1 := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	x2 := []string{"Diego", "Alberto", "Lopez", "Chaverra"}

	fmt.Println(x1)
	sort.Ints(x1)
	fmt.Println(x1)
	fmt.Println(x2)
	sort.Strings(x2)
	fmt.Println(x2)
}
