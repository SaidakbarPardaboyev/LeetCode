/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	var backtracking func(root *TreeNode, targetSum int, pathSum int) bool
	backtracking = func(root *TreeNode, targetSum int, pathSum int) bool {
		if root == nil {
			return false
		}
		pathSum += root.Val

		if root.Left == nil && root.Right == nil {
			return targetSum == pathSum
		}

		if backtracking(root.Left, targetSum, pathSum) {
			return true
		}
		if backtracking(root.Right, targetSum, pathSum) {
			return true
		}

		return false
	}

	return backtracking(root, targetSum, 0)
}