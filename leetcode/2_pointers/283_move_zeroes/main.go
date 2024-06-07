package main

func moveZeroes(nums []int) {
	index := 0

	for _, num := range nums {
		if num == 0 {
			continue
		}
		nums[index] = num
		index++
	}

	for index < len(nums) {
		nums[index] = 0
		index++
	}
}
