package main

import "testing"

func TestCalculateNeededGoals(t *testing.T) {
	tests := []struct {
		round1    *Round
		round2    *Round
		placement int
		want      int
	}{
		// Case 1: Clear aggregate victory for team 1
		{&Round{Team1Score: 2, Team2Score: 1}, &Round{Team1Score: 1, Team2Score: 0}, HOME, 0},
		// Case 2: Aggregate tied, team 1 need 1 goal to win
		{&Round{Team1Score: 1, Team2Score: 2}, &Round{Team1Score: 2, Team2Score: 1}, HOME, 1},
		// Case 2: Aggregate tied, team 1 wins on away goals
		{&Round{Team1Score: 1, Team2Score: 2}, &Round{Team1Score: 2, Team2Score: 1}, AWAY, 1},
		// Case 3: Aggregate and away goals tied
		{&Round{Team1Score: 1, Team2Score: 2}, &Round{Team1Score: 2, Team2Score: 1}, AWAY, 1},
		// Case 4: Team 2 needs to overcome deficit, home scenario
		{&Round{Team1Score: 1, Team2Score: 2}, &Round{Team1Score: 1, Team2Score: 2}, HOME, 3},
		// Case 5: Team 2 needs to overcome deficit, away scenario
		{&Round{Team1Score: 1, Team2Score: 2}, &Round{Team1Score: 1, Team2Score: 2}, AWAY, 2},
		// Case 5: Team 2 needs to overcome deficit, away scenario | win by away goals
		{&Round{Team1Score: 1, Team2Score: 1}, &Round{Team1Score: 2, Team2Score: 2}, AWAY, 0},
		// Case 5: Team 2 needs to overcome deficit, home scenario | win by total goals
		{&Round{Team1Score: 1, Team2Score: 1}, &Round{Team1Score: 2, Team2Score: 2}, HOME, 1},
		// Case 6: test1 from the task description
		{&Round{Team1Score: 0, Team2Score: 0}, &Round{Team1Score: 0, Team2Score: 0}, HOME, 1},
		// Case 6: test2 from the task description
		{&Round{Team1Score: 0, Team2Score: 2}, &Round{Team1Score: 0, Team2Score: 3}, AWAY, 5},
		// Case 6: test3 from the task description
		{&Round{Team1Score: 0, Team2Score: 2}, &Round{Team1Score: 0, Team2Score: 3}, HOME, 6},
		{&Round{Team1Score: 2, Team2Score: 1}, &Round{Team1Score: 1, Team2Score: 0}, HOME, 0},
		{&Round{Team1Score: 1, Team2Score: 2}, &Round{Team1Score: 2, Team2Score: 1}, HOME, 1},

		{&Round{Team1Score: 1, Team2Score: 2}, &Round{Team1Score: 2, Team2Score: 1}, AWAY, 1},
		{&Round{Team1Score: 2, Team2Score: 3}, &Round{Team1Score: 3, Team2Score: 2}, HOME, 1},
		{&Round{Team1Score: 0, Team2Score: 1}, &Round{Team1Score: 1, Team2Score: 0}, AWAY, 1},
		{&Round{Team1Score: 1, Team2Score: 2}, &Round{Team1Score: 0, Team2Score: 1}, HOME, 3},
		{&Round{Team1Score: 1, Team2Score: 3}, &Round{Team1Score: 0, Team2Score: 1}, AWAY, 4},
		{&Round{Team1Score: 5, Team2Score: 5}, &Round{Team1Score: 5, Team2Score: 5}, HOME, 1},
		{&Round{Team1Score: 0, Team2Score: 2}, &Round{Team1Score: 2, Team2Score: 0}, AWAY, 1},
		{&Round{Team1Score: 4, Team2Score: 2}, &Round{Team1Score: 0, Team2Score: 3}, AWAY, 2},
		{&Round{Team1Score: 4, Team2Score: 3}, &Round{Team1Score: 0, Team2Score: 3}, HOME, 2},
	}

	for _, tt := range tests {
		got := CalculateNeededGoals(tt.round1, tt.round2, tt.placement)
		if got != tt.want {
			t.Errorf("CalculateNeededGoals(%v, %v, %d) = %d, want %d", tt.round1, tt.round2, tt.placement, got, tt.want)
		}
	}
}
