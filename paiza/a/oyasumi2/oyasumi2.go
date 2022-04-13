package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		days int
		holidayArray []int
		dayCount int = 0
		whiteCheckArray []bool
	)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Scan(&days)

	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	holidayArray = convIntArray(input)

	for i := 0; i < len(holidayArray); i++ {
		whiteCheckArray = append(whiteCheckArray, whiteCheck(holidayArray[i:]))
	}

	for i := 0; i < len(whiteCheckArray); i++ {
		if i == 0 || i == (len(whiteCheckArray) -1 ) || dayCount == 0{
			if whiteCheckArray[i] {
				dayCount++
			}
		} else {
			if whiteCheckArray[i] && whiteCheckArray[i-1] {
				dayCount++
			}
		}
	}
	dayCount += 3

	if dayCount == 3 {
		dayCount = 0
	}

	if dayCount > days {
		dayCount = days
	}

	fmt.Println(dayCount)

}

func convIntArray(strArray []string) (intSlices []int) {
	for _, v := range strArray {
		x, _ := strconv.Atoi(v)
		intSlices = append(intSlices, x)
	}

	return intSlices
}

func whiteCheck(intArray []int) bool {
	x := 0
	if len(intArray) < 7 {
		for i := 0; i < len(intArray); i++ {
			if intArray[i] == 0 {
				x++
			}
		}
	} else {
		for i := 0; i < 7; i++ {
			if intArray[i] == 0 {
				x++
			}
		}
	}

	if x >= 2 {
		return true
	}
	return false
}