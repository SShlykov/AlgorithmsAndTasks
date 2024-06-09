package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(num int) int {
	return -1
}

func guessNumber(n int) int {
	l, r := 0, n+1

	for r-l > 1 {
		m := (r + l) / 2
		switch guess(m) {
		case 0:
			return m
		case 1:
			l = m
		case -1:
			r = m
		}
	}

	return -1
}
