package main

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return true
		}

		if nums[l] == nums[m] && nums[r] == nums[m] {
			l++
			r--
		} else if nums[l] <= nums[m] {
			if nums[l] <= target && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[m] < target && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

	return false
}

// O(n)
func searchOn(nums []int, target int) bool {
	if len(nums) == 1 {
		return nums[0] == target
	}
	if searchOn(nums[:len(nums)/2], target) || searchOn(nums[len(nums)/2:], target) {
		return true
	}
	return false
}
