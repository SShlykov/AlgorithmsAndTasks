package three_sum

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		x        []int
		expected [][]int
	}{
		{"2 slices", []int{-1, 0, 1, 2, -1, -4}, [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		}},
		{"empty solutions", []int{0, 1, 1}, [][]int{}},
		{"single variant", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := threeSum(tc.x)
			if len(result) != len(tc.expected) {
				t.Errorf("threeSum(%#v) = %#v, expected %#v", tc.x, result, tc.expected)
			}
		})
	}
}
