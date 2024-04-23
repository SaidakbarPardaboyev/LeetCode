/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	rightSideNodes := []int{}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		rightSideNodes = append(rightSideNodes, queue[length-1].Val)
		for i := 0; i < length; i++ {
			currentNode := queue[0]
			queue = queue[1:]

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}

			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
	}

	return rightSideNodes
}