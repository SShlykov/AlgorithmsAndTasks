package powx_n

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		x        int
		expected int
	}{
		{"целое", 4, 2},
		{"не целое", 8, 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := mySqrt(tc.x)
			if result != tc.expected {
				t.Errorf("mySqrt(%#v) = %v, expected %v", tc.x, result, tc.expected)
			}
		})
	}
}
