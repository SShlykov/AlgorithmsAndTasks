package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])

	l, r := 0, n*m

	good := func(index int) bool {
		i, j := index/n, index%n
		return matrix[i][j] <= target
	}

	for r-l > 1 {
		mid := (r + l) / 2
		if good(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	i, j := l/n, l%n

	return matrix[i][j] == target
}
