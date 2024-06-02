package zigzag_conversion

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		n1       string
		n2       int
		expected string
	}{
		{"123", "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"123", "PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"123", "A", 1, "A"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := convert(tc.n1, tc.n2)
			if result != tc.expected {
				t.Errorf("convert(%#v, %#v) = %v, expected %v", tc.n1, tc.n2, result, tc.expected)
			}
		})
	}
}
