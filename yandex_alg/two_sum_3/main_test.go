package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	name          string
	arr           []int
	sum           int
	expectedValue []int
}

func Test_twoSum(t *testing.T) {
	tests := []testCase{
		{name: "1", arr: []int{-1, -1, -9, -7, 3, -6}, sum: 2, expectedValue: []int{3, -1}},
		{name: "2", arr: []int{-3, 1, 1, 2, 6, 6, 8, 10}, sum: 100, expectedValue: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := twoSum(tt.arr, tt.sum)
			fmt.Println(res)
		})
	}
}
