package string_to_integer_atoi

const limitMax = 2147483647
const limitMin = -2147483648

func myAtoi(s string) int {
	var i, num int
	sign := 1
	if s == "" {
		return 0
	}
	for i < len(s) && s[i] == 32 {
		i++
	}

	if i >= len(s) {
		return 0
	}

	if s[i] == 43 {
		i++
	} else if s[i] == 45 {
		sign = -1
		i++
	}
	for ; i < len(s); i++ {
		if s[i] != 0 && (s[i] < 48 || s[i] > 57) {
			return num * sign
		}
		n := int(s[i])
		num = num*10 + (n - 48)
		if num*sign < limitMin {
			return limitMin
		} else if num*sign > limitMax {
			return limitMax
		}
	}

	return num * sign
}
