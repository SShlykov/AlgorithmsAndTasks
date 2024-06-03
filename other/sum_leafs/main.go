package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func sumRec(n *Node) int {
	if n == nil {
		return 0
	}

	if n.left == nil && n.right == nil {
		return n.val
	}

	return sumRec(n.left) + sumRec(n.right)
}

func sumIterative(root *Node) int {
	if root == nil {
		return 0
	}

	sum := 0
	stack := []*Node{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.left == nil && node.right == nil {
			sum += node.val
		}

		if node.right != nil {
			stack = append(stack, node.right)
		}

		if node.left != nil {
			stack = append(stack, node.left)
		}
	}

	return sum
}

func main() {
	root := &Node{
		val: 1,
		left: &Node{
			val: 2,
			left: &Node{
				val: 4,
			},
			right: &Node{
				val: 5,
			},
		},
		right: &Node{
			val: 3,
			left: &Node{
				val: 6,
			},
			right: &Node{
				val: 7,
			},
		},
	}

	fmt.Println(sumRec(root))       // Ожидаем 22
	fmt.Println(sumIterative(root)) // Ожидаем 22
}
