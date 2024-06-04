package uorder

// uorder([]int{1,1,0,0,0,1,1,1}) -> 3
// uorder([]int{}) -> 0
// uorder([]int{0,0,0}) -> 0
// uorder([]int{1,1,1,1,1}) -> 5

func uorder(list []int) int {
	if len(list) == 0 {
		return 0
	}
	curr, best := 0, 0

	for _, val := range list {
		if val > 0 {
			curr++
			best = max(best, curr)
		} else {
			curr = 0
		}
	}
	return best
}
