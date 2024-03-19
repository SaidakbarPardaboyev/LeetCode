func longestOnes(nums []int, k int) int {
	res := 0
	summa := 0
	start := 0
	for end := 0; end < len(nums); end++ {
		summa += nums[end]

		if end-start+1-summa > k {
			res = max(res, end-start)
			summa -= nums[start]
			start++
		}
	}
	res = max(res, len(nums)-start)

	return res
}