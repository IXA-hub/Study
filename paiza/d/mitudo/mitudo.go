package main

import (
	"fmt"
	"os"
)

func main() {
	var x, y int

	if _, err := fmt.Scan(&x); err != nil {
		os.Exit(1)
	}
	if _, err := fmt.Scan(&y); err != nil {
		os.Exit(1)
	}

	fmt.Printf("%d\n", newFunc(x, y))

}

func newFunc(x, y int) int {
	return x*y
}