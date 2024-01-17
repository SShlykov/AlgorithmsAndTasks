package merge_strings_alternately

func mergeAlternately(word1 string, word2 string) (str string) {
	l1, l2 := len(word1), len(word2)
	maxLen := l1
	if maxLen < l2 {
		maxLen = l2
	}

	for i := 0; i < maxLen; i++ {
		var letter1, letter2 string
		if i < l1 {
			letter1 = string(word1[i])
		}
		if i < l2 {
			letter2 = string(word2[i])
		}
		str = str + letter1 + letter2
	}

	return
}
