package main

import (
	"fmt"
	"testing"
)

func Test_movingAverage(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	window := 4

	res := movingAverage(arr, window)
	fmt.Println(res)
}
