package container_with_most_water

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		n1       []int
		expected int
	}{
		{"123", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"123", []int{1, 1}, 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxArea(tc.n1)
			if result != tc.expected {
				t.Errorf("maxArea(%#v) = %v, expected %v", tc.n1, result, tc.expected)
			}
		})
	}
}
