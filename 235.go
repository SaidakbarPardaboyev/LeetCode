/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	res, _ := FindCommanAncestor(root, p, q)
	return res
}

func FindCommanAncestor(cur, target1, target2 *TreeNode) (*TreeNode, int) {
	if cur == nil {
		return nil, 0
	}

	res := 0
	if cur == target1 {
		res++
	}
	if cur == target2 {
		res++
	}

	anc, tem := FindCommanAncestor(cur.Left, target1, target2)
	res += tem
	if res == 2 && anc == nil {
		return cur, res
	} else if res == 2 && anc != nil {
		return anc, res
	}
	anc2, tem := FindCommanAncestor(cur.Right, target1, target2)
	res += tem
	if res == 2 && anc2 == nil {
		return cur, res
	} else if res == 2 && anc2 != nil {
		return anc2, res
	}

	return nil, res
}