package merge_k_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var tmp *ListNode
	l := len(lists)
	if l == 0 {
		return tmp
	}

	res := lists

	for len(res) > 1 {
		var newRes []*ListNode
		i := 0
		for ; i+1 < len(res); i += 2 {
			merge(res[i], res[i+1], &newRes)
		}

		if i < len(res) {
			newRes = append(newRes, res[i])
		}

		res = newRes
	}

	if len(res) == 0 {
		return nil
	}

	return res[0]
}

func merge(l1, l2 *ListNode, res *[]*ListNode) {
	var head = ListNode{}
	cur := &head
	for l1 != nil || l2 != nil {
		if l1 != nil && (l2 == nil || l1.Val < l2.Val) {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}

	if head.Next != nil {
		*res = append(*res, head.Next)
	}
}
