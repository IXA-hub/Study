package main

import (
	"fmt"
)

func main() {
	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}

	for _, v := range numbers {
		fmt.Println(block_stairs(&v))
	}
	fmt.Println(numbers)
}

func block_stairs(number *int) int{
	blocks := 1
	for ;;blocks++ {
		*number -= blocks
		if *number <= blocks {
			break
		}
	}
	return blocks
}