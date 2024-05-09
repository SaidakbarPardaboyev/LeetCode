func maximumHappinessSum(happiness []int, k int) int64 {
	MergeSort(happiness, 0, len(happiness)-1)

	maximumHappiness := 0
	happinessDecreasingBy := 0

	for i := 0; i < len(happiness) && k > 0; i++ {
		if happiness[i]-happinessDecreasingBy >= 0 {
			maximumHappiness += happiness[i] - happinessDecreasingBy
		}
		happinessDecreasingBy++
		k--
	}

	return int64(maximumHappiness)
}

func MergeSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	MergeSort(nums, start, mid)
	MergeSort(nums, mid+1, end)

	Merge(nums, start, mid, end)
}

func Merge(nums []int, start, mid, end int) {
	left := make([]int, mid-start+1)
	copy(left, nums[start:mid+1])
	right := make([]int, end-mid)
	copy(right, nums[mid+1:end+1])

	l := 0
	r := 0
	n := start

	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			nums[n] = left[l]
			l++
		} else {
			nums[n] = right[r]
			r++
		}
		n++
	}

	for l < len(left) {
		nums[n] = left[l]
		n++
		l++
	}
	for r < len(right) {
		nums[n] = right[r]
		n++
		r++
	}
}