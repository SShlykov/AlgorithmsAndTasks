package powx_n

import (
	"math"
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		x        float64
		n        int
		expected float64
	}{
		{"целое", 2.00000, 10, 1024.00000},
		{"не целое", 2.10000, 3, 9.26100},
		{"отрицательное", 2.00000, -2, 0.25},
	}

	arc := math.Pow(10, 6)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := myPow(tc.x, tc.n)
			if math.Floor(result*arc)/arc != math.Floor(tc.expected*arc)/arc {
				t.Errorf("myPow(%#v, %#v) = %v, expected %v", tc.x, tc.n, result, tc.expected)
			}
		})
	}
}
