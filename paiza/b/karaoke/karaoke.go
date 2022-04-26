package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var (
		N, M  int
		music []int
		input [][]int
		score []int
	)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Scan(&N, &M)

	for i := 0; i < N; i++ {
		score = append(score, 100)
	}

	for i := 0; i < M; i++ {
		scanner.Scan()
		hoge, _ := strconv.Atoi(scanner.Text())
		music = append(music, hoge)
	}

	for i := 0; i < N; i++ {
		var fuga []int
		for j := 0; j < M; j++ {
			scanner.Scan()
			hoge, _ := strconv.Atoi(scanner.Text())
			fuga = append(fuga, hoge)
		}
		input = append(input, fuga)
	}

	for i := 0; i < N; i++ {

		for j := 0; j < M; j++ {
			x := music[j]
			x -= input[i][j]
			if x < 0 {
				x = x*-1
			}
			switch {
			case x > 0 && x <= 10:
				score[i] -= 1
			case x > 10 && x <= 20:
				score[i] -= 2
			case x > 20 && x <= 30:
				score[i] -= 3
			case x > 30:
				score[i] -= 5
			}
		}
	}

	for i := 0; i < len(score); i++ {
		if score[i] < 0 {
			score[i] = 0
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(score)))
	fmt.Println(score[0])
}