package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	switch {
	case list1 == nil:
		return list2
	case list2 == nil:
		return list1
	case list1.Val <= list2.Val:
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	case list2.Val < list1.Val:
		list2.Next = mergeTwoLists(list2.Next, list1)
		return list2
	default:
		panic("unexpected case")
	}
}
