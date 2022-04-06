package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	arr := []int{20, 25, 17, 37, 45, 77, 100, 1, 57}

	for _, v := range arr {
		convSeconds(v)
	}
}