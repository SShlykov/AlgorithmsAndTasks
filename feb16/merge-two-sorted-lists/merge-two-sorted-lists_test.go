package merge_two_sorted_lists

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{
			name:  "both lists empty",
			list1: nil,
			list2: nil,
			want:  nil,
		},
		{
			name:  "list1 empty, list2 not empty",
			list1: nil,
			list2: &ListNode{Val: 1, Next: nil},
			want:  &ListNode{Val: 1, Next: nil},
		},
		{
			name:  "list1 not empty, list2 empty",
			list1: &ListNode{Val: 1, Next: nil},
			list2: nil,
			want:  &ListNode{Val: 1, Next: nil},
		},
		{
			name:  "both lists have one element",
			list1: &ListNode{Val: 1, Next: nil},
			list2: &ListNode{Val: 2, Next: nil},
			want:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
		},
		{
			name:  "lists with elements in increasing order",
			list1: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}},
			list2: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}},
			want:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}},
		},
		// Add more test cases as necessary
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.list1, tt.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
