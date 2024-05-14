/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	nums1 := []int{}
	nums2 := []int{}

	nums1 = dfs(root1)
	nums2 = dfs(root2)

	return Merge(nums1, nums2)
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	nums := []int{}
	nums = append(nums, dfs(root.Left)...)
	nums = append(nums, root.Val)
	nums = append(nums, dfs(root.Right)...)
	return nums
}

func Merge(nums1, nums2 []int) []int {
	res := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	for i < len(nums1) {
		res = append(res, nums1[i])
		i++
	}

	for j < len(nums2) {
		res = append(res, nums2[j])
		j++
	}
	return res
}