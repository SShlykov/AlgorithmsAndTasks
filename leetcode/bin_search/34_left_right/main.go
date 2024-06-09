package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 8))
}

//		 | 0   1   2   3   4   5   6 |
//		 | 5,  7,  7,  8,  8,  10    |
//	     | t   t   t   f   f   f     | l
//	     | f   f   f   f   f   t     | r
//
// l = -1; r = 6 -> m = 3 -> nums[3] = 8
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{searchL(nums, target), searchR(nums, target)}
}

func searchL(nums []int, target int) int {
	l, r := -1, len(nums)-1

	for r-l > 1 {
		m := (r + l) / 2
		if goodL(m, target, nums) {
			l = m
		} else {
			r = m
		}
	}

	if nums[r] == target {
		return r
	} else {
		return -1
	}
}

func goodL(curr, target int, nums []int) bool {
	return nums[curr] < target
}

func searchR(nums []int, target int) int {
	l, r := 0, len(nums)

	for r-l > 1 {
		m := (r + l) / 2
		if goodR(m, target, nums) {
			r = m
		} else {
			l = m
		}
	}

	if nums[l] == target {
		return l
	} else {
		return -1
	}
}

func goodR(curr, target int, nums []int) bool {
	return nums[curr] > target
}
