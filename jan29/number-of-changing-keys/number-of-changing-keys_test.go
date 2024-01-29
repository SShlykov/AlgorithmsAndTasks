package number_of_changing_keys

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		x        string
		expected int
	}{
		{"целое", "aAbBcC", 2},
		{"без изменений", "AaAaAaaA", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := countKeyChanges(tc.x)
			if result != tc.expected {
				t.Errorf("mySqrt(%#v) = %v, expected %v", tc.x, result, tc.expected)
			}
		})
	}
}
