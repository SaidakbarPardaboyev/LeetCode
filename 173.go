/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	nums []int
	ind  int
}

func Constructor(root *TreeNode) BSTIterator {
	nums := []int{}

	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}

		helper(root.Left)
		nums = append(nums, root.Val)
		helper(root.Right)
	}
	helper(root)

	BSTIterator := BSTIterator{
		nums: nums,
		ind:  0,
	}
	return BSTIterator
}

func (this *BSTIterator) Next() int {
	res := this.nums[this.ind]
	this.ind++
	return res
}

func (this *BSTIterator) HasNext() bool {
	return this.ind < len(this.nums)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */