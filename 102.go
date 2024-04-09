/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	levelOrder := [][]int{}

	queue := []*TreeNode{root}
	levelOrder = append(levelOrder, []int{root.Val})

	for len(queue) > 0 {
		l := len(queue)
		curLevel := []int{}
		for i := 0; i < l; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
				curLevel = append(curLevel, cur.Left.Val)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
				curLevel = append(curLevel, cur.Right.Val)
			}
		}
		if len(curLevel) > 0 {
			levelOrder = append(levelOrder, curLevel)
		}
	}
	return levelOrder
}