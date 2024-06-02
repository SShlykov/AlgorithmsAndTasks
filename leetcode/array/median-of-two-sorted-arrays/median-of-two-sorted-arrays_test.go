package median_of_two_sorted_arrays

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		n1       []int
		n2       []int
		expected float64
	}{
		{"123", []int{1, 3}, []int{2}, 2},
		{"123", []int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findMedianSortedArrays(tc.n1, tc.n2)
			if result != tc.expected {
				t.Errorf("findMedianSortedArrays(%#v, %#v) = %v, expected %v", tc.n1, tc.n2, result, tc.expected)
			}
		})
	}
}
