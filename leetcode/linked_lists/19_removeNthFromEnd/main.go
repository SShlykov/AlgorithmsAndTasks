package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	fast := dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	slow := dummy
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
