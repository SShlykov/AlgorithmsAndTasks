package string_to_integer_atoi

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"1", "42", 42},
		{"1", "   -42", -42},
		{"1", "4193 with words", 4193},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := myAtoi(tc.input)
			if result != tc.expected {
				t.Errorf("myAtoi(%s) = %d, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
