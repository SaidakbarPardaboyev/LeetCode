import "math"

func jump(nums []int) int {
	minimumJampsPerIndex := []int{0}

	for i := len(nums) - 2; i >= 0; i-- {
		// To find minimum jamp to the end from the index
		minJumpFromCurrentIndex := math.MaxInt64
		for j := 0; j < nums[i] && j < len(minimumJampsPerIndex); j++ {
			if minimumJampsPerIndex[j] > -1 && minimumJampsPerIndex[j] < minJumpFromCurrentIndex {
				minJumpFromCurrentIndex = minimumJampsPerIndex[j]
			}
		}

		// check if nums[i] == 0 add -1 to list, if there is no number in list add 1,
		// else add to the list with min jumb plas
		if minJumpFromCurrentIndex < math.MaxInt64 {
			minimumJampsPerIndex = append([]int{minJumpFromCurrentIndex + 1}, minimumJampsPerIndex...)
		} else if len(minimumJampsPerIndex) == 1 {
			minimumJampsPerIndex = append([]int{1}, minimumJampsPerIndex...)
		} else {
			minimumJampsPerIndex = append([]int{-1}, minimumJampsPerIndex...)
		}
	}

	return minimumJampsPerIndex[0]
}