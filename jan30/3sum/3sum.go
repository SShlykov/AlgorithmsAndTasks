package three_sum

import "sort"

type List []int

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := v + nums[l] + nums[r]
			if sum > 0 {
				r -= 1
			} else if sum < 0 {
				l += 1
			} else {
				res = append(res, []int{v, nums[l], nums[r]})
				l += 1
				for nums[l] == nums[l-1] && l < r {
					l += 1
				}
			}
		}
	}

	return res
}
