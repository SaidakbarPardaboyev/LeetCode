/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	return depthFirstSearch(root, 0, targetSum)
}

func depthFirstSearch(node *TreeNode, sum, targetSum int) bool {
	if node == nil {
		return false
	}

	sum += node.Val

	if node.Left == nil && node.Right == nil {
		return sum == targetSum
	}

	return depthFirstSearch(node.Left, sum, targetSum) || depthFirstSearch(node.Right, sum, targetSum)
}