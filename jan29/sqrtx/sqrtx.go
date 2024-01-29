package powx_n

func mySqrt(x int) int {
	i := 1
	for i*i < x {
		i += 1
	}
	if i*i == x {
		return i
	}
	return i - 1
}

func mySqrtSlower(x int) int {
	var list []int
	for i := 0; i*i <= x; i++ {
		list = append(list, i)
	}

	return list[len(list)-1]
}
