package main

import "testing"

// TestIsPalindrome используется для тестирования функции isPalindrome.
func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"PositiveCase", 121, true},
		{"NegativeCase", 123, false},
		{"SingleDigit", 7, true},
		{"Zero", 0, true},
		{"NegativeNumber", -121, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isPalindrome(tc.input)
			if result != tc.expected {
				t.Errorf("isPalindrome(%d) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
