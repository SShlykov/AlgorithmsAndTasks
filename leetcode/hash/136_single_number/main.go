package main

func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i] // <- bitwise xor operation ( x ^ x = 1; 1 ^ x = x)
	}
	return res
}
