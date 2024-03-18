func productExceptSelf(nums []int) []int {
	product := 1
	zero := 0
	for i := range nums {
		if nums[i] != 0 {
			product *= nums[i]
		} else {
			zero++
		}
	}

	for i := range nums {
		if nums[i] != 0 {
			if zero > 0 {
				nums[i] = 0
			} else {
				nums[i] = product / nums[i]
			}
		} else {
			if zero == 1 {
				nums[i] = product
			} else {
				nums[i] = 0
			}
		}
	}
	return nums
}