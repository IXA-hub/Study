package main

import (
	"fmt"
)

func main() {
	var (
		gondora_count, group_count int
		gondora_capacity []int
		group_menber []int
		x = 0
	)

	fmt.Scan(&gondora_count, &group_count)

	answer := make([]int, gondora_count)

	for i := 0; i < gondora_count; i++ {
		x := 0
		fmt.Scan(&x)
		gondora_capacity = append(gondora_capacity, x)
	}

	for i := 0; i < group_count; i++ {
		x := 0
		fmt.Scan(&x)
		group_menber = append(group_menber, x)
	}


	for i := 0; x < group_count; i++{
		if i == gondora_count {
			i = 0
		}

		if group_menber[x] <= gondora_capacity[i] {
			answer[i] += group_menber[x]
			x++
		} else {
			answer[i] += gondora_capacity[i]
			group_menber[x] -= gondora_capacity[i]
		}
	}

	for _, v := range answer {
		fmt.Printf("%d\n", v)
	}

}