package merge_k_sorted_lists

import (
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name  string
		lists []*ListNode
		want  *ListNode
	}{
		{
			name:  "all lists are empty",
			lists: []*ListNode{nil, nil},
			want:  nil,
		},
		{
			name:  "single list in the array",
			lists: []*ListNode{{Val: 1, Next: &ListNode{Val: 2, Next: nil}}},
			want:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
		},
		{
			name:  "multiple lists, all empty",
			lists: []*ListNode{nil, nil, nil},
			want:  nil,
		},
		{
			name:  "multiple lists, some empty",
			lists: []*ListNode{nil, {Val: 1, Next: nil}, nil},
			want:  &ListNode{Val: 1, Next: nil},
		},
		{
			name: "multiple lists with elements in increasing order",
			lists: []*ListNode{
				{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 7, Next: nil}}},
				{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 8, Next: nil}}},
				{Val: 3, Next: &ListNode{Val: 6, Next: &ListNode{Val: 9, Next: nil}}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: nil}}}}}}}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
