package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1,2,3,4,5
// fast + 2 -> 3 -> 5 -> br
// slow + 1 -> 2 -> 3
// if next == nil || next.next == nil -> break

// 1,2,3,4,5,6
// fast + 2 -> 3 -> 5 -> br
// slow + 1 -> 2 -> 3
// if next == nil || next.next == nil -> break

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == nil || fast.Next == nil {
			break
		}
	}

	return slow
}
