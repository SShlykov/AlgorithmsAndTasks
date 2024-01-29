package regular_expression_matching

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		p        string
		expected bool
	}{
		{"falsy", "aa", "a", false},
		{"truthy", "aa", "a*", true},
		{"truthy", "ab", ".*", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isMatch(tc.s, tc.p)
			if result != tc.expected {
				t.Errorf("isMatch(%#v, %#v) = %v, expected %v", tc.s, tc.p, result, tc.expected)
			}
		})
	}
}
