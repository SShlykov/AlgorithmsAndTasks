package reverse_integer

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"123", 123, 321},
		{"123", -123, -321},
		{"123", 120, 21},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := reverse(tc.input)
			if result != tc.expected {
				t.Errorf("reverse(%d) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
