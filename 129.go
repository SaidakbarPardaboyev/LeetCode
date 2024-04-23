/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	return dfsTofindSumOfRootToLeaf(root, 0)
}

func dfsTofindSumOfRootToLeaf(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}

	sum = sum*10 + node.Val

	if node.Left == nil && node.Right == nil {
		return sum
	}
	return dfsTofindSumOfRootToLeaf(node.Left, sum) + dfsTofindSumOfRootToLeaf(node.Right, sum)
	return sum
}