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
		respondent, question int
		answer               [][]string
		average              []int
	)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Scan(&respondent, &question)

	for i := 0; i < respondent; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		answer = append(answer, input)
	}

	for i := 0; i < question; i++ {
		validCount := 0
		var validAnswer []int
		validTotal := 0
		for j := 0; j < respondent; j++ {
			x, y := inputCheck(answer[j][i]); if y == true && x >= 0 && x <= 100 {
				validAnswer = append(validAnswer, x)
				validCount++
			}
		}
		if validCount != 0 {
			for _, v := range validAnswer {
				validTotal += v
			}
			average = append(average, validTotal/validCount)
		} else {
			average = append(average, 0)
		}
	}
	
	for _, v := range average {
		fmt.Printf("%d\n", v)
	}
}

func inputCheck(input string) (int, bool) {
	i, err := strconv.Atoi(input)
	if err != nil {
		return i, false
	}
	return i, true
}
