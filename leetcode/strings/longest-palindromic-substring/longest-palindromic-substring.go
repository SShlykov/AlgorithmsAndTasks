package longest_palindromic_substring

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLength := 0, 1
	for i := 0; i < len(s); i++ {
		left, right := extendPalindrome(s, i, i)
		length := right - left + 1
		if length > maxLength {
			maxLength = length
			start = left
		}

		left, right = extendPalindrome(s, i, i+1)
		length = right - left + 1
		if length > maxLength {
			maxLength = length
			start = left
		}
	}

	return s[start : start+maxLength]
}

func extendPalindrome(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
