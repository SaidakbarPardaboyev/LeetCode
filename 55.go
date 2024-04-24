func canJump(nums []int) bool {
	canJumbToPerIndexs := []bool{true}

	for i := len(nums) - 2; i >= 0; i-- {
		// check if there is a path from current index to the end
		isTherePath := false
		for j := 0; j < len(canJumbToPerIndexs) && j < nums[i]; j++ {
			if canJumbToPerIndexs[j] {
				isTherePath = true
				break
			}
		}

		// if there is a path, add true to the from of the list
		// because by this true you will know you can to go the end
		// by visiting the current index
		if isTherePath {
			canJumbToPerIndexs = append([]bool{true}, canJumbToPerIndexs...)
		} else {
			canJumbToPerIndexs = append([]bool{false}, canJumbToPerIndexs...)
		}
	}

	return canJumbToPerIndexs[0]
}