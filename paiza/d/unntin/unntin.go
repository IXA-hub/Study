package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stdin := scanner.Text()

	kukan, err := strconv.Atoi(stdin); if err != nil {
		fmt.Printf("not int")
		os.Exit(1)
	}

	if !(1<=kukan) || !(kukan<=30) {
		fmt.Printf("out range")
	}

	ryoukin := newFunc(kukan)
	fmt.Printf("%d\n", ryoukin)
}

func newFunc(i int) int {
	return 100 + i*10
}