func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt
	cur := 0
	start := 0
	for end := 0; end < len(nums); end++ {
		cur += nums[end]
		for cur >= target {
			res = min(res, end-start+1)
			cur -= nums[start]
			start++
		}
	}
	if res > len(nums) {
		return 0
	}
	return res
}