/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findSmallestNode(root *TreeNode) int {
	curNode := root
	for curNode != nil && curNode.Left != nil {
		curNode = curNode.Left
	}
	return curNode.Val
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			root.Val = findSmallestNode(root.Right)
			root.Right = deleteNode(root.Right, root.Val)
		}
	}

	return root
}