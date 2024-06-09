package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	getMinNodeIdx := func(list []*ListNode) int {
		minIndex := -1
		minValue := 1 << 32
		for i := 0; i < len(list); i++ {
			if list[i] == nil {
				continue
			}

			if list[i].Val < minValue {
				minIndex = i
				minValue = list[i].Val
			}
		}

		return minIndex
	}

	dummy := &ListNode{}
	curr := dummy

	for {
		minNodeIndex := getMinNodeIdx(lists)
		if minNodeIndex == -1 {
			curr.Next = nil
			break
		}

		curr.Next = lists[minNodeIndex]
		curr = curr.Next
		lists[minNodeIndex] = lists[minNodeIndex].Next
	}

	return dummy.Next
}
