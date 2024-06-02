package majnumbers

func majorityElement(nums []int) int {
	var (
		candidate = 0
		count     = 0
	)

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if candidate == num {
			count++
		} else {
			count--
		}
	}

	return candidate
}
