package main

func reverseVector(vector []int) []int {
	if len(vector) < 2 {
		return vector
	}

	left, right := 0, len(vector)-1

	for left < right {
		vector[left], vector[right] = vector[right], vector[left]
		left++
		right--
	}

	return vector
}
