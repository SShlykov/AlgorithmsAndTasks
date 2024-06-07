package main

import "sort"

func meetingRooms(meetings [][]int) int {
	if len(meetings) < 2 {
		return len(meetings)
	}
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	return 0
}
