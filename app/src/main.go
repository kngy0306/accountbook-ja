package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
	}
	// [170 165 150 160 175 180 155]
	fmt.Println("===========")

	sort.Ints(t)
	for i := 0; i < x; i++ {
		fmt.Println(t[i])
	}
}
