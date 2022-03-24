package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	s := stdin.Text()

	split_stdin := strings.Split(s, " ")
	newList := make([]int, len(split_stdin))

	newList = convStringListToIntList(split_stdin)

	count := 0

	for {
		//スライスnewListの中に偶数があるかどうかをチェックする
		if !isEvenNumberList(newList) {
			break
		}
		//スライスnewListを1/2にする
		newList = divideByTwo(newList)
		count++
	}
	fmt.Println(count)
}

func convStringListToIntList(list []string) []int {
	newList := make([]int, len(list))
	for i, v := range list {
		a, _ := strconv.Atoi(v)
		newList[i] = a
	}
	return newList
}

func isEvenNumberList(list []int) bool {
	result := true
	for _, v := range list {
		if v%2 != 0 {
			result = false
			break
		}
	}
	return result
}

func divideByTwo(list []int) []int {
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = v / 2
	}
	return newList
}
