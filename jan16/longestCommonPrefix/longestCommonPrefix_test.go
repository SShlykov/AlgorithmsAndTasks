package longestCommonPrefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected string
	}{
		{"fl_test", []string{"flower", "flow", "flight"}, "fl"},
		{"empty", []string{"dog", "racecar", "car"}, ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := longestCommonPrefix(tc.input)
			if result != tc.expected {
				t.Errorf("longestCommonPrefix(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
