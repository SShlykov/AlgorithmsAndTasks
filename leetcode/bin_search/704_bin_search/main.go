package main

// time - O(log(n))
// memory O(1)
func search(nums []int, target int) int {
	l, r := 0, len(nums)
	for r-l > 1 {
		marker := (l + r) / 2
		if good(nums[marker], target) {
			l = marker
		} else {
			r = marker
		}
	}

	if nums[l] == target {
		return l
	} else {
		return -1
	}
}

func good(curr, target int) bool {
	return curr <= target
}
