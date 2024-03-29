/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	queue1 := []*TreeNode{p}
	queue2 := []*TreeNode{q}
	if p.Val != q.Val {
		return false
	}

	for len(queue1) > 0 {
		for i := 0; i < len(queue1); i++ {
			curNode1 := queue1[0]
			curNode2 := queue2[0]
			ls1 := []int{}
			ls2 := []int{}
			// fmt.Println(curNode1.Val)
			// fmt.Println(curNode2.Val)

			queue1 = queue1[1:]
			queue2 = queue2[1:]
			lamp := 0
			lamp2 := 0

			if curNode1.Left != nil {
				queue1 = append(queue1, curNode1.Left)
				lamp++
				ls1 = append(ls1, curNode1.Left.Val)
			}
			if curNode1.Right != nil {
				queue1 = append(queue1, curNode1.Right)
				lamp2++
				ls1 = append(ls1, curNode1.Right.Val)
			}

			if curNode2.Left != nil {
				queue2 = append(queue2, curNode2.Left)
				lamp++
				ls2 = append(ls2, curNode2.Left.Val)
			}
			if curNode2.Right != nil {
				queue2 = append(queue2, curNode2.Right)
				lamp2++
				ls2 = append(ls2, curNode2.Right.Val)
			}

			if !(lamp == 0 || lamp == 2) || !(lamp2 == 0 || lamp2 == 2) {
				return false
			} else {
				for i := range ls1 {
					if ls1[i] != ls2[i] {
						return false
					}
				}
			}
		}
	}
	return true
}