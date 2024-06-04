package anagram

func SetFromString(s string) map[rune]int {
	set := make(map[rune]int)

	for _, letter := range []rune(s) {
		set[letter]++
	}
	return set
}

func Anagram(a, b string) bool {
	aSet := SetFromString(a)
	bSet := SetFromString(b)
	if len(aSet) != len(bSet) {
		return false
	}

	for letter, count := range aSet {
		if count != bSet[letter] {
			return false
		}
	}

	return true
}
