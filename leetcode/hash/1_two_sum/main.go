package main

func twoSum(nums []int, target int) []int {
	set := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		val := nums[i]

		if ind, ok := set[target-val]; ok {
			return []int{ind, i}
		} else {
			set[val] = i
		}
	}

	return []int{}
}
