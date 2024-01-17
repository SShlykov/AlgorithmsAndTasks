package merge_strings_alternately

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		input2   string
		expected string
	}{
		{"1", "abc", "pqr", "apbqcr"},
		{"1", "ab", "pqrs", "apbqrs"},
		{"1", "abcd", "pq", "apbqcd"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := mergeAlternately(tc.input, tc.input2)
			if result != tc.expected {
				t.Errorf("mergeAlternately(%s, %s) = %v, expected %v", tc.input, tc.input2, result, tc.expected)
			}
		})
	}
}
