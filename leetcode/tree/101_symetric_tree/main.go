package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return symetric(root.Left, root.Right)
}

func symetric(node1, node2 *TreeNode) bool {
	switch {
	case node1 == nil || node2 == nil:
		return node1 == node2
	case node1.Val == node2.Val:
		return symetric(node1.Left, node2.Right) && symetric(node1.Right, node2.Left)
	default:
		return false
	}
}
