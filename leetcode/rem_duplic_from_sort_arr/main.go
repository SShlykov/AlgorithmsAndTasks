package main

func removeDuplicates(nums []int) int {
	k := 0
	prev := -101

	for _, num := range nums {
		if num != prev {
			nums[k] = num
			k += 1
			prev = num
			continue
		}
	}
	nums = nums[:k]

	return k
}
