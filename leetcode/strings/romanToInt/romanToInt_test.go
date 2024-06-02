package romanToInt

import "testing"

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"III", "III", 3},
		{"LVIII", "LVIII", 58},
		{"MCMXCIV", "MCMXCIV", 1994},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := romanToInt(tc.input)
			if result != tc.expected {
				t.Errorf("romanToInt(%s) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
