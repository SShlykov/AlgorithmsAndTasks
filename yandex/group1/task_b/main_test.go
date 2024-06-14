package main

import "testing"

func TestCalculateNeededGoals(t *testing.T) {
	tests := []struct {
		round1    *Round
		round2    *Round
		placement int
		want      int
	}{
		{&Round{2, 1}, &Round{1, 0}, HOME, 0},
		{&Round{1, 2}, &Round{2, 1}, HOME, 1},
		{&Round{1, 2}, &Round{2, 1}, AWAY, 1},
		{&Round{1, 2}, &Round{2, 1}, AWAY, 1},
		{&Round{1, 2}, &Round{1, 2}, HOME, 3},
		{&Round{1, 2}, &Round{1, 2}, AWAY, 2},
		{&Round{1, 1}, &Round{2, 2}, AWAY, 0},
		{&Round{1, 1}, &Round{2, 2}, HOME, 1},
		{&Round{0, 0}, &Round{0, 0}, HOME, 1},
		{&Round{0, 2}, &Round{0, 3}, AWAY, 5},
		{&Round{0, 2}, &Round{0, 3}, HOME, 6},
		{&Round{2, 1}, &Round{1, 0}, HOME, 0},
		{&Round{1, 2}, &Round{2, 1}, HOME, 1},
		{&Round{1, 2}, &Round{2, 1}, AWAY, 1},
		{&Round{2, 3}, &Round{3, 2}, HOME, 1},
		{&Round{0, 1}, &Round{1, 0}, AWAY, 1},
		{&Round{1, 2}, &Round{0, 1}, HOME, 3},
		{&Round{1, 3}, &Round{0, 1}, AWAY, 4},
		{&Round{5, 5}, &Round{5, 5}, HOME, 1},
		{&Round{0, 2}, &Round{2, 0}, AWAY, 1},
		{&Round{4, 2}, &Round{0, 3}, AWAY, 2},
		{&Round{4, 3}, &Round{0, 3}, HOME, 2},
	}

	for _, tt := range tests {
		got := CalculateNeededGoals(tt.round1, tt.round2, tt.placement)
		if got != tt.want {
			t.Errorf("CalculateNeededGoals(%v, %v, %d) = %d, want %d", tt.round1, tt.round2, tt.placement, got, tt.want)
		}
	}
}
