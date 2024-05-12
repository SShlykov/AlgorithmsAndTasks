package containsDuplicate

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})

	for _, num := range nums {
		if _, exist := set[num]; exist {
			return true
		}

		set[num] = struct{}{}
	}
	return false
}
