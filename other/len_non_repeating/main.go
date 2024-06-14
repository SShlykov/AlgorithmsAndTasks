package main

func isReapeted(set map[int]int) bool {
	for _, count := range set {
		if count > 1 {
			return true
		}
	}
	return false
}

func lenMaxNonRepeatingElements(list []int) int {
	if len(list) < 2 {
		return len(list)
	}

	maxWindow := 0
	set := make(map[int]int)

	for i := 0; i < len(list); i++ {
		set[list[i]]++

		if isReapeted(set) {
			set[list[i-maxWindow]]--
		} else {
			maxWindow++
		}

	}

	return maxWindow
}

func GetMaxLen(list []int) int {
	set := make(map[int]int)
	prev, best := 0, 0

	for i := 0; i < len(list); i++ {
		prev = max(prev, set[list[i]])
		best = max(best, i-prev+1)
		set[list[i]] = i + 1
	}

	return best
}
