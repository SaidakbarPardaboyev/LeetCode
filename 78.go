func subsets(nums []int) [][]int {
	var subsets [][]int
	var curSubset []int
	FindAllSubsets(0, nums, &subsets, curSubset)
	return subsets
}

func FindAllSubsets(ind int, nums []int, subsets *[][]int, curSubset []int) {
	if ind >= len(nums) {
		*subsets = append(*subsets, append([]int{}, curSubset...))
		return
	}

	curSubset = append(curSubset, nums[ind])
	FindAllSubsets(ind+1, nums, subsets, curSubset)

	curSubset = curSubset[:len(curSubset)-1]
	FindAllSubsets(ind+1, nums, subsets, curSubset)
}