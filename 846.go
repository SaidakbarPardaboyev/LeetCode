func isNStraightHand(hand []int, groupSize int) bool {
	nums := FindCountOfNumbers(hand)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	start := 0
	for start < len(nums) {
		prev, end, count := nums[start][0], start+1, 1
		nums[start][1]--
		for end < len(nums) && count < groupSize {
			for end < len(nums) && nums[end][1] == 0 {
				end++
			}

			if end >= len(nums) || nums[end][0]-prev > 1 {
				return false
			}

			count++
			nums[end][1]--
			prev = nums[end][0]
			end++
		}

		if count != groupSize {
			return false
		}

		for start < len(nums) && nums[start][1] == 0 {
			start++
		}
	}

	return true
}

func FindCountOfNumbers(nums []int) [][]int {
	numsWithCount := map[int]int{}

	for _, num := range nums {
		numsWithCount[num]++
	}

	res := [][]int{}
	for num, count := range numsWithCount {
		res = append(res, []int{num, count})
	}

	return res
}