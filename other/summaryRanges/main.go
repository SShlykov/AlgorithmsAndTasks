package main

import "fmt"

func summary(nums []int) []string {
	l, r := 0, 0
	result := make([]string, 0)

	for l < len(nums) {
		for r+1 < len(nums) && nums[r]+1 == nums[r+1] {
			r++
		}

		if r != l {
			result = append(result, fmt.Sprintf("%d->%d", nums[l], nums[r]))
		} else {
			result = append(result, fmt.Sprintf("%d", nums[l]))
		}

		l++
		r++
	}

	return result
}
