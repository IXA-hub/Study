package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	inputStr := scanner.Text()

	s, err := strconv.Atoi(inputStr); if err != nil {
		os.Exit(1)
	}

	if !(1 <= s) || !(s <= 100) {
		os.Exit(1)
	}
	fmt.Printf("%d\n", convSeconds(s))

}

func convSeconds(i int) int {
	return i * 60
}