package main

func isPerfectSquare(num int) bool {
	good := func(x int) bool {
		return x*x > num
	}

	l, r := 0, num+1

	for r-l > 1 {
		m := (l + r) / 2
		if good(m) {
			r = m
		} else {
			l = m
		}
	}

	return l*l == num
}
