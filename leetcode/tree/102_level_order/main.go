package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int

	traverse(root, 0, &res)
	return res
}

func traverse(node *TreeNode, level int, list *[][]int) {
	if node == nil {
		return
	}
	if len(*list) <= level {
		*list = append(*list, []int{})
	}
	(*list)[level] = append((*list)[level], node.Val)
	traverse(node.Left, level+1, list)
	traverse(node.Right, level+1, list)
}
