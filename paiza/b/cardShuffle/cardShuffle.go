package main

import (
	"fmt"
)

func main() {
	var (
		N, M, K int
		cards []int
		after_Shuffle []int
	)

	fmt.Scan(&N, &M, &K)
	
	for i := 1; i <= N; i++ {
		cards = append(cards, i)
	}

	x := 1
	for i := 0; i < M; i++ {
		y[i] := cards[M*i:M*x]
		x++
	}
	fmt.Println(y)
	fmt.Print(after_Shuffle)
}

func funcA(cards []int, x, y int) []int {
	return cards[x:y]
}