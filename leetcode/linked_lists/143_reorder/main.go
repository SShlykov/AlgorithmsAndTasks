package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func preMiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	tmp := preMiddleNode(head)
	p2 := reverseList(tmp.Next)

	tmp.Next = nil

	p1 := head
	newHead := p1

	for p2 != nil {
		p1Next := p1.Next
		p1.Next = p2
		p1 = p2
		p2 = p1Next
	}

	head = newHead
}
