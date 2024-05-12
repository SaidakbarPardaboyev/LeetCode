func combine(n int, k int) [][]int {
	combinations := [][]int{}

	var helper func(number, n int, curCombination []int)
	helper = func(number, n int, curCombination []int) {
		if len(curCombination) >= k {
			combinations = append(combinations, append([]int{}, curCombination...))
			return
		}

		for i := number; i <= n; i++ {
			curCombination = append(curCombination, i)
			helper(i+1, n, curCombination)
			curCombination = curCombination[:len(curCombination)-1]
		}
	}
	helper(1, n, []int{})
	return combinations
}