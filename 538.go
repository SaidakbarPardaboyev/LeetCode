/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var helper func(root *TreeNode)

	helper = func(root *TreeNode) {
		if root == nil {
			return
		}

		helper(root.Right)
		sum += root.Val
		root.Val = sum
		helper(root.Left)
	}
	helper(root)
	return root
}