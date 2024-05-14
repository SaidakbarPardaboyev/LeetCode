func getMaximumGold(nums [][]int) int {
	resMax := 0
	visited := map[int]bool{}

	var helper func(i, j, curMax int)

	helper = func(i, j, curMax int) {
		if _, ok := visited[i*len(nums[0])+j]; ok {
			return
		}
		curMax += nums[i][j]
		resMax = max(resMax, curMax)
		visited[i*len(nums[0])+j] = true

		if j+1 < len(nums[0]) && nums[i][j+1] != 0 {
			helper(i, j+1, curMax)
		}
		if i+1 < len(nums) && nums[i+1][j] != 0 {
			helper(i+1, j, curMax)
		}
		if i-1 >= 0 && nums[i-1][j] != 0 {
			helper(i-1, j, curMax)
		}
		if j-1 >= 0 && nums[i][j-1] != 0 {
			helper(i, j-1, curMax)
		}
		delete(visited, i*len(nums[0])+j)
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			helper(i, j, 0)
		}
	}
	return resMax
}