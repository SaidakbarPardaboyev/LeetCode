func countAlternatingSubarrays(nums []int) int64 {
	numberOfSubarray := 1
	sizeOfLastSabarray := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			sizeOfLastSabarray += 1
		} else {
			sizeOfLastSabarray = 1
		}
		numberOfSubarray += sizeOfLastSabarray
	}

	return int64(numberOfSubarray)
}