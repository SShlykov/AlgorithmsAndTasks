package validParentheses

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"1", "()", true},
		{"2", "()[]{}", true},
		{"3", "(]", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isValid(tc.input)
			if result != tc.expected {
				t.Errorf("IsValid(%s) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
