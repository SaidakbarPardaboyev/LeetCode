import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	return inOrderTravel(root, &k)
}

func inOrderTravel(node *TreeNode, k *int) int {
	res := 0
	if node == nil {
		return 0
	}
	res = max(res, inOrderTravel(node.Left, k))
	*k--
	if *k == 0 {
		return node.Val
	}
	fmt.Println(node.Val)
	res = max(res, inOrderTravel(node.Right, k))
	return res
}