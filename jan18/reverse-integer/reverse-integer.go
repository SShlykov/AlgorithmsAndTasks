package reverse_integer

const limitMax = 2147483647
const limitMin = -2147483647

func reverse(x int) (reversed int) {
	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	if reversed > limitMax || reversed < limitMin {
		return 0
	}
	return
}
