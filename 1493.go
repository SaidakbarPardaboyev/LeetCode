func longestSubarray(nums []int) int {
	summa := 0
	start := 0
	res := 0
	for end := 0; end < len(nums); end++ {
		summa += nums[end]
		if end-start+1-summa <= 1 {
			res = max(res, summa)
		} else {
			summa -= nums[start]
			start++
		}
	}

	if summa == len(nums) {
		return summa - 1
	}
	return res
}