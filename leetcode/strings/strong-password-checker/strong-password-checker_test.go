package strong_password_checker

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"1", "a", 5},
		{"2", "aA1", 3},
		{"3", "1337C0d3", 0},
		{"4", "bbaaaaaaaaaaaaaaacccccc", 8},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := strongPasswordChecker(tc.input)
			if result != tc.expected {
				t.Errorf("strongPasswordChecker(%s) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
