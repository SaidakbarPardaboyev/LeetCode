func combinationSum3(k int, n int) [][]int {
	res := [][]int{}

	var FindAllCombinations func(num int, curCom []int, curSum int)
	FindAllCombinations = func(num int, curCom []int, curSum int) {
		if len(curCom) > k || curSum > n {
			return
		}
		if len(curCom) == k && curSum == n {
			res = append(res, append([]int{}, curCom...))
			return
		}

		for i := num; i <= 9; i++ {
			curCom = append(curCom, i)
			FindAllCombinations(i+1, curCom, curSum+i)
			curCom = curCom[:len(curCom)-1]
		}
	}
	FindAllCombinations(1, []int{}, 0)
	return res
}