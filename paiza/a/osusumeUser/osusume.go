package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

type user struct {
	id             int
	frend          []int
	popular     int
	recommendation []int
}

func main() {
	var (
		userCount, friendShipCount int
		userInfo                   []user
	)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Scan(&userCount, &friendShipCount)

	scanner.Scan()
	popularity := strings.Split(scanner.Text(), " ")
	intPopularity := convIntArray(popularity)

	for i := 0; i < userCount; i++ {
		userInfo = append(userInfo, user{id: i+1, popular: intPopularity[i]})
	}

	for i := 0; i < friendShipCount; i++ {
		scanner.Scan()
		fuga := strings.Split(scanner.Text(), " ")
		hoge := convIntArray(fuga)
		userInfo[hoge[0]-1].frend = append(userInfo[hoge[0]-1].frend, hoge[1])
		userInfo[hoge[1]-1].frend = append(userInfo[hoge[1]-1].frend, hoge[0])
	}

	for i:= 0; i < len(userInfo); i++ {
		fmt.Println(userRecommendation(userInfo[i].frend, userInfo))
		//fmt.Println(popularMatchUser(userRecommendation(userInfo[i].frend, userInfo), userInfo))
	}

}

func convIntArray(strArray []string) (intSlices []int) {
	for _, v := range strArray {
		x, _ := strconv.Atoi(v)
		intSlices = append(intSlices, x)
	}

	return intSlices
}

func userRecommendation(frend []int, users []user) int {
	var x []int
	for i := 0; i < len(users); i++ {
		for _, v := range frend {
			if users[i].id != v {
				x = append(x, users[i].popular)
			}
		}
	}

	fmt.Println(x)

	sort.Ints(sort.IntSlice(x))
	return x[len(x) - 1]
}

func popularMatchUser(popular int, users []user) int {
	for _, v := range users {
		if popular == v.popular {
			return v.id
		}
	}
	return -1
}