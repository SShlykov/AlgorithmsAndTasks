package greatest_common_divisor_of_strings

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		input2   string
		expected string
	}{
		{"1", "ABCABC", "ABC", "ABC"},
		{"1", "ABABAB", "ABAB", "AB"},
		{"1", "LEET", "CODE", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := gcdOfStrings(tc.input, tc.input2)
			if result != tc.expected {
				t.Errorf("gcdOfStrings(%s, %s) = %v, expected %v", tc.input, tc.input2, result, tc.expected)
			}
		})
	}
}
