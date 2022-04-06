package main

import (
	"testing"
	"fmt"
)

type testcase struct{
	Case []int
	Answer int
}

func TestNewFunc(t *testing.T) {
	m1 := map[int]testcase{
		1: {Case: []int{4, 15}, Answer: 60},
		2: {Case: []int{3, 12}, Answer: 36},
	}

	for _, v := range m1 {
		x := newFunc(v.Case[0], v.Case[1])
		if x != v.Answer {
			fmt.Printf("fail test:%d is not %d", x, v.Answer)
			t.Errorf("test fail")
		}
	}
}