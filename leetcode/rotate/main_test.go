package rotate

import "testing"

func Test_rotate(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6, 7}
	round := 3

	rotate(l, round)
}
