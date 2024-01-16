package romanToInt

import "strings"

func romanToInt(s string) int {
	strList := strings.Split(s, "")
	res := 0
	length := len(strList)
	lastValue := 0

	for i := length - 1; i >= 0; i-- {
		value := convertLetterToInt(strList[i])
		if value < lastValue {
			res -= value
		} else {
			res += value
		}
		lastValue = value
	}

	return res
}

func convertLetterToInt(letter string) (num int) {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	return romanMap[letter]
}
