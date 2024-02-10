func maximumTripletValue(nums []int) int64 {
	res := int64(0)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				tem := int64((nums[i] - nums[j]) * nums[k])
				if tem > res {
					res = tem
				}
			}
		}
	}
	return res
}