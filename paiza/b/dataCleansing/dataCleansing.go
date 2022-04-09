package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	var (
		respondent, question int
		answer [][]string
	)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Scan(&respondent, &question)

	for i := 0; i < respondent; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		answer = append(answer, input)
	}

	for i := 0; i < len(answer); i++ {
		for j := 0; j < question; j++ {
			fmt.Print(answer[j][i])
		}
		fmt.Println(i)
	}

	fmt.Print(answer)
}

func inputCheck(input []string) bool {
	check := true
	label1:
		for _, v := range input {
			_, err := strconv.Atoi(v); if err != nil {
				check = false
				break label1
			}
		}
	return check
}