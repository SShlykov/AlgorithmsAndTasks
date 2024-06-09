package main

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)
	for r-l > 1 {
		mark := (l + r) / 2
		if good(mark, arr) {
			l = mark
		} else {
			r = mark
		}
	}

	return l
}

func good(curr int, arr []int) bool {
	return arr[curr-1] < arr[curr]
}
