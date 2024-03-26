/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			curNode := queue[0]
			queue = queue[1:]
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
		count++
	}
	return count
}