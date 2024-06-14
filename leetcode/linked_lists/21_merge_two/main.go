package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	getVal := func(list *ListNode) int {
		if list == nil {
			return 1 << 31
		}
		return list.Val
	}

	curr := &ListNode{}
	dummy := curr
	for list1 != nil || list2 != nil {
		if getVal(list1) <= getVal(list2) {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	return dummy.Next
}
