package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	arr := []int{5, 7, 10, 13, 17, 21}

	for _, v := range arr {
		if !(newFunc(v) == 100 + (10*v)) {
			t.Errorf("fail test")
		}
	}
}