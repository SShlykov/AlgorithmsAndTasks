package quick_sort

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		n1       []int
		expected []int
	}{
		{"123", []int{1, 3, 2}, []int{1, 2, 3}},
		{"123", []int{1, 3, 4, 2}, []int{1, 2, 3, 4}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := quickSort(tc.n1)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("quickSort(%#v)= %#v, expected %#v", tc.n1, result, tc.expected)
			}
		})
	}
}
