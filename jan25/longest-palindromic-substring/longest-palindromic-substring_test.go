package longest_palindromic_substring

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		n1       string
		expected string
	}{
		{"one", "babad", "bab"},
		{"two", "cbbd", "bb"},
		{"ccc", "ccc", "ccc"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestPalindrome(tc.n1)
			if result != tc.expected {
				t.Errorf("longestPalindrome(%#v) = %v, expected %v", tc.n1, result, tc.expected)
			}
		})
	}
}
