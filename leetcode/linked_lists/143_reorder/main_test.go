package main

import (
	"fmt"
	"testing"
)

func Test_reorder(t *testing.T) {
	head := &ListNode{Val: 1}
	v1 := &ListNode{Val: 2}
	v2 := &ListNode{Val: 3}
	v3 := &ListNode{Val: 4}
	v2.Next = v3
	v1.Next = v2
	head.Next = v1

	reorderList(head)
	for {
		if head == nil {
			break
		}
		fmt.Println(head.Val)
		head = head.Next
	}
}
