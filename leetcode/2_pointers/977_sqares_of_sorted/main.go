package main

import "slices"

func sortedSquares(nums []int) []int {
	var result []int
	l, r := 0, len(nums)-1

	for l <= r {
		if abs(nums[l]) > abs(nums[r]) {
			result = append(result, nums[l]*nums[l])
			l++
		} else {
			result = append(result, nums[r]*nums[r])
			r--
		}
	}
	slices.Reverse(result)

	return result
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
