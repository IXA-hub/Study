package main

import (
	"fmt"
)

func main () {
	var (
		x string
	)

	fmt.Scan(&x)

	printPlus(x)
	fmt.Printf("+%s+\n", x)
	printPlus(x)
}

func printPlus(s string) {
	for i := 0; i <= len(s); i++ {
		fmt.Printf("+")
	}
	fmt.Printf("+\n")
}