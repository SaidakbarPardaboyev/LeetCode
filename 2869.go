func minOperations(nums []int, k int) int {
	res := []int{}
	sch := 0
	for i := len(nums) - 1; i >= 0; i-- {
		sch += 1
		if nums[i] <= k {
			lamp := true
			for _, v := range res {
				if v == nums[i] {
					lamp = false
					break
				}
			}
			if lamp {
				res = append(res, nums[i])
				if len(res) == k {
					return sch
				}
			}
		}
	}
	return 0
}