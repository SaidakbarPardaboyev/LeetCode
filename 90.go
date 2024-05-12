func subsetsWithDup(nums []int) [][]int {
	subsets := [][]int{}
	curSubset := []int{}

	MergeSort(nums, 0, len(nums)-1)
	FindSubsetsWithDublicates(0, nums, &subsets, curSubset)
	return subsets
}

func MergeSort(nums []int, start, end int) {
	if !(start < end) {
		return
	}

	mid := (start + end) / 2
	MergeSort(nums, start, mid)
	MergeSort(nums, mid+1, end)

	Merge(nums, start, mid, end)
}

func Merge(nums []int, start, mid, end int) {
	left := append([]int{}, nums[start:mid+1]...)
	right := append([]int{}, nums[mid+1:end+1]...)

	l, r := 0, 0
	n := start
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
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
		l++
		n++
	}

	for r < len(right) {
		nums[n] = right[r]
		n++
		r++
	}
}

func FindSubsetsWithDublicates(ind int, nums []int, subsets *[][]int, curSubset []int) {
	if ind >= len(nums) {
		*subsets = append(*subsets, append([]int{}, curSubset...))
		return
	}

	curSubset = append(curSubset, nums[ind])
	FindSubsetsWithDublicates(ind+1, nums, subsets, curSubset)

	curSubset = curSubset[:len(curSubset)-1]
	for ind+1 < len(nums) && nums[ind+1] == nums[ind] {
		ind++
	}
	FindSubsetsWithDublicates(ind+1, nums, subsets, curSubset)
}

