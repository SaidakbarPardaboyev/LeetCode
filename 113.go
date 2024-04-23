/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	return depthFirstSearch(root, []int{}, targetSum, 0)
}

func depthFirstSearch(node *TreeNode, res []int, targetSum, currentSum int) [][]int {
	if node == nil {
		return [][]int{}
	}

	res = append(res, node.Val)
	currentSum += node.Val

	if node.Left == nil && node.Right == nil && currentSum == targetSum {
		return [][]int{res}
	}

	// Without copying it to another recursion It would give a pointer of the list
	newGroupOfNodes := make([]int, len(res))
	copy(newGroupOfNodes, res)

	return append(depthFirstSearch(node.Left, newGroupOfNodes, targetSum, currentSum), depthFirstSearch(node.Right, newGroupOfNodes, targetSum, currentSum)...)
}