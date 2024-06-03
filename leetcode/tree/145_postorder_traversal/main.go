package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var res []int
	traverse(root, &res)
	return res
}

func traverse(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	traverse(root.Left, res)
	traverse(root.Right, res)
	*res = append(*res, root.Val)
}
