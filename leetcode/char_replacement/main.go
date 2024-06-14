package char_replacement

func characterReplacement(s string, k int) int {
	m := map[byte]int{}
	maxLen := 0
	maxRepeat := byte(0)
	start := 0
	end := 0

	for end < len(s) {
		char := s[end]
		m[char]++
		count := m[char]

		if maxRepeat == 0 || m[maxRepeat] < count {
			maxRepeat = char
		}

		if end-start+1-m[maxRepeat] > k {
			m[s[start]]--
			start++
		}

		if maxLen < end-start+1 {
			maxLen = end - start + 1
		}
		end++
	}

	return maxLen
}
