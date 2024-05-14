/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	var root *TreeNode
	var helper func(root *TreeNode, val int) *TreeNode

	helper = func(root *TreeNode, val int) *TreeNode {
		if root == nil {
			return &TreeNode{
				Val: val,
			}
		}

		if root.Val > val {
			root.Left = helper(root.Left, val)
		} else {
			root.Right = helper(root.Right, val)
		}

		return root
	}

	for _, num := range preorder {
		root = helper(root, num)
	}

	return root
}