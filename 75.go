func sortColors(nums []int) {
	res := QuickSort(nums)
	for i := 0; i < len(nums); i++ {
		nums[i] = res[i]
	}
}

func QuickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[len(nums)/2]
	less := []int{}
	greater := []int{}

	for i := 0; i < len(nums); i++ {
		if i == len(nums)/2 {
			continue
		}

		if nums[i] <= pivot {
			less = append(less, nums[i])
		} else {
			greater = append(greater, nums[i])
		}
	}

	less = QuickSort(less)
	greater = QuickSort(greater)
	less = append(less, pivot)
	less = append(less, greater...)
	return less
}