package rotate

func rotate(nums []int, k int) {
	if k == 0 || k == len(nums) {
		return
	}
	if k > len(nums) {
		k = k % len(nums)
	}
	reverse(nums[:len(nums)-k])
	reverse(nums[len(nums)-k:])
	reverse(nums)
}

func reverse(nums []int) {
	if len(nums) == 0 {
		return
	}
	end := len(nums) - 1
	mid := len(nums) / 2
	for i := 0; i < mid; i++ {
		nums[i], nums[end-i] = nums[end-i], nums[i]
	}
}
