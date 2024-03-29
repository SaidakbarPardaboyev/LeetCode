package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintBinaryTree(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	curNode := root
	for curNode != nil || len(stack) > 0 {
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}

		curNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curNode.Val)

		curNode = curNode.Right
	}
	return res
}

func isValidBST(root *TreeNode) bool {
	ls := PrintBinaryTree(root)
	if len(ls) < 2 {
		return true
	}
	for i := 0; i < len(ls)-1; i++ {
		if !(ls[i] < ls[i+1]) {
			return false
		}
	}
	return true
}

func main() {
	node1 := &TreeNode{Val: 5}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 7}
	node4 := &TreeNode{Val: 1}
	node5 := &TreeNode{Val: 3}
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 8}
	node8 := &TreeNode{Val: 4}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	node5.Right = node8

	fmt.Println(isValidBST(node1))
}
