func removeDuplicates(nums []int) int {
	count := 0
	start := 0
	prevNum := math.MinInt64
	k := 0
	for i := 0; i < len(nums); i++ {
		if prevNum != nums[i] {
			count = 0
			prevNum = nums[i]
		}
		count++
		if count > 2 {
			k++
			continue
		}
		nums[start] = prevNum
		start++
	}
	return len(nums) - k
}