/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	ls1 := ReturnLeaves(root1)
	ls2 := ReturnLeaves(root2)
	if len(ls1) != len(ls2) {
		return false
	}
	for i := 0; i < len(ls1); i++ {
		if ls1[i] != ls2[i] {
			return false
		}
	}
	return true
}

func ReturnLeaves(root *TreeNode) []int {
	leaves := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		lamp := 0
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
			lamp++
		}
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
			lamp++
		}
		if lamp == 0 {
			leaves = append(leaves, curNode.Val)
		}
	}
	return leaves
}