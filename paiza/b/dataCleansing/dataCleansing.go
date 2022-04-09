package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"reflect"
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

	for i := 0; i < len(answer); i++ {
		validCount := 0
		validAnswer := []int{}
		validTotal := 0
		for j := 0; j < question; j++ {
			x, y := inputCheck(answer[i][j]); if y == true {
				validAnswer = append(validAnswer, x)
				validCount++
			}
		}
		println(validAnswer)
		for _, v := range validAnswer {
			validTotal += v
		}

		average = append(average, validTotal/validCount)
	}
	fmt.Print(average)
}

func inputCheck(input string) (int, bool) {
	i, err := strconv.Atoi(input)
	if err != nil {
		return i, false
	}
	return i, true
}
