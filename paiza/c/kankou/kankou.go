package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		N, K, answer   int
		nArray, kArray []int
		move           [][]int
	)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		nArray = append(nArray, n)
	}

	for i := 0; i < N; i++ {
		scanner.Scan()
		hoge := strings.Split(scanner.Text(), " ")
		move = append(move, convIntArray(hoge))
	}

	fmt.Scan(&K)

	for i := 0; i < K; i++ {
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		kArray = append(kArray, k)
	}

	for i := 0; i < len(kArray); i++ {
		answer += nArray[kArray[i]-1]
	}

	for i := 0; i < len(kArray)-1; i++ {
		answer += move[kArray[i]-1][kArray[i+1]-1]
	}

	fmt.Println(answer)
}

func convIntArray(hoge []string) []int {
	var fuga []int
	for _, v := range hoge {
		x, _ := strconv.Atoi(v)
		fuga = append(fuga, x)
	}
	return fuga
}
