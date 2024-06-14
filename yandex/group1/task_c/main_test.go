package main

import "testing"

func TestGetCountOperations(t *testing.T) {
	tests := []struct {
		name               string
		countNeededSpaces  int
		expectedOperations int
	}{
		{
			name:               "No operations needed",
			countNeededSpaces:  0,
			expectedOperations: 0,
		},
		{
			name:               "Less than one TAB, only spaces",
			countNeededSpaces:  3,
			expectedOperations: 2, // 1 TAB, 1 backspace
		},
		{
			name:               "Exactly one TAB, prefer TAB over spaces",
			countNeededSpaces:  4,
			expectedOperations: 1, // 1 TAB
		},
		{
			name:               "More than one TAB, no exact divisor",
			countNeededSpaces:  10,
			expectedOperations: 4, // 2 TAB, 2 spaces
		},
		{
			name:               "Large number, testing efficiency",
			countNeededSpaces:  100,
			expectedOperations: 25, // 25 TAB, 0 spaces
		},
		// Add more cases as necessary, especially edge cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCountOperations(tt.countNeededSpaces); got != tt.expectedOperations {
				t.Errorf("getCountOperations(%d) = %d, want %d", tt.countNeededSpaces, got, tt.expectedOperations)
			}
		})
	}
}
