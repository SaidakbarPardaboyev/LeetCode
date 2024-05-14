func combinationSum(nums []int, target int) [][]int {
	resCombinations := [][]int{}
	var backtracking func(curCombination []int, sum, ind int)

	backtracking = func(curCombination []int, sum, ind int) {
		if ind >= len(nums) || sum > target {
			return
		}

		if sum == target {
			resCombinations = append(resCombinations, append([]int{}, curCombination...))
			return
		}
		curCombination = append(curCombination, nums[ind])
		backtracking(curCombination, sum+nums[ind], ind)

		curCombination = curCombination[:len(curCombination)-1]
		backtracking(curCombination, sum, ind+1)
	}
	backtracking([]int{}, 0, 0)

	return resCombinations
}