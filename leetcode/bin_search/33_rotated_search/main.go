package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	fmt.Println(findOffset(nums))
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	offset := findOffset(nums)
	l, r := 0, len(nums)

	good := func(curr int) bool {
		return nums[curr] <= target
	}

	for r-l > 1 {
		m := (r + l) / 2

		if good((m + offset) % len(nums)) {
			l = m
		} else {
			r = m
		}
	}

	realLeft := (l + offset) % len(nums)

	if nums[realLeft] == target {
		return realLeft
	} else {
		return -1
	}
}

func findOffset(nums []int) int {
	l, r := -1, len(nums)-1
	good := func(ind int) bool {
		return nums[ind] > nums[len(nums)-1]
	}

	for r-l > 1 {
		m := (r + l) / 2
		if good(m) {
			l = m
		} else {
			r = m
		}
	}

	return r
}
