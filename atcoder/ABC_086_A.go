package main

import (
	"fmt"
)

func main() {
	var (
		a, b int
	)

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	if (a+b)%2 == 1 {
		fmt.Println("Odd")
		return
	}

	fmt.Println("Even")
}
