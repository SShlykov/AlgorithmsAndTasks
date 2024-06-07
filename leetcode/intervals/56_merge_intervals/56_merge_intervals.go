package main

import (
	"fmt"
	"sort"
)

const (
	X = iota
	Y
)

func crossed(a []int, b []int) bool {
	return max(a[X], b[X]) <= min(a[Y], b[Y])
}

func doMerge(a []int, b []int) []int {
	return []int{min(a[X], b[X]), max(a[Y], b[Y])}
}

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{intervals[0]}

	for _, curr := range intervals {
		i := len(result) - 1
		if crossed(result[i], curr) {
			result[i] = doMerge(result[i], curr)
		} else {
			result = append(result, curr)
		}
	}

	return result
}
func main() {
	int := [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}

	fmt.Println(merge(int))
}
