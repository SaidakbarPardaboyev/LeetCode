func permute(nums []int) [][]int {

	var FindPermutations func(i int) [][]int
	FindPermutations = func(i int) [][]int {
		if i == len(nums) {
			return [][]int{{}}
		}
		resPer := [][]int{}
		curPer := FindPermutations(i + 1)
		for _, per := range curPer {
			for j := 0; j < len(per)+1; j++ {
				newPer := append([]int{}, per[:j]...)
				newPer = append(newPer, nums[i])
				newPer = append(newPer, per[j:]...)
				resPer = append(resPer, newPer)
			}
		}
		return resPer
	}
	return FindPermutations(0)
}