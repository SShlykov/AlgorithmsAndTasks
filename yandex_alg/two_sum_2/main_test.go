package main

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	arr := []int{-1, -1, -9, -7, 3, -6}
	sum := 2

	res := twoSum(arr, sum)
	fmt.Println(res)
}
