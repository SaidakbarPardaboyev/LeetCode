import "sort"

func minOperationsToMakeMedianK(nums []int, k int) int64 {
	sort.Ints(nums)
	// fmt.Println(nums)
	// the number of opeartion to make the middle
	var numberOfOparation int64
	numberOfOparation += int64(nums[len(nums)/2] - k)
	if numberOfOparation < 0 {
		numberOfOparation = -numberOfOparation
	}
	// nums[len(nums)/2] = k

	i := len(nums)/2 - 1
	middle := k

	for i >= 0 && nums[i] > middle {
		numberOfOparation += int64(nums[i] - k)
		// nums[i] = k
		i--
	}

	i = len(nums)/2 + 1

	for i < len(nums) && nums[i] < middle {
		numberOfOparation += int64(k - nums[i])
		// nums[i] = k
		i++
	}

	// fmt.Println(nums)
	return numberOfOparation
}