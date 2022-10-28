package main

import (
	"fmt"
)

func factoring(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factoring(n-1)
	}
}

func main() {
	var num int
	fmt.Printf("input a positive whole number:")
	fmt.Scanln(&num)

	fmt.Println(factoring(num))
}
