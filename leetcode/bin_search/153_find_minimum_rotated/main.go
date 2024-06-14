package main

func findMin(nums []int) int {
	offset := findOffset(nums)

	return nums[offset]
}

func findOffset(nums []int) int {
	last := nums[len(nums)-1]
	good := func(curr int) bool {
		return curr > last
	}

	l, r := -1, len(nums)

	for r-l > 1 {
		m := (l + r) / 2
		if good(nums[m]) {
			l = m
		} else {
			r = m
		}
	}

	return r
}
