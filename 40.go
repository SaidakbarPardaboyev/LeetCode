func combinationSum2(candidates []int, target int) [][]int {
	combinations := [][]int{}

	MergeSort(candidates, 0, len(candidates)-1)

	var helper func(ind, comSum int, curCombination []int)
	helper = func(ind, comSum int, curCombination []int) {
		if comSum == target {
			combinations = append(combinations, append([]int{}, curCombination...))
			return
		}
		if comSum > target || ind >= len(candidates) {
			return
		}

		curCombination = append(curCombination, candidates[ind])
		helper(ind+1, comSum+candidates[ind], curCombination)

		for ind+1 < len(candidates) && candidates[ind] == candidates[ind+1] {
			ind++
		}
		curCombination = curCombination[:len(curCombination)-1]
		helper(ind+1, comSum, curCombination)
	}
	helper(0, 0, []int{})

	return combinations
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
		r++
		n++
	}
}