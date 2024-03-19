func findDifference(nums1 []int, nums2 []int) [][]int {
	dic1 := make(map[int]int)
	dic2 := make(map[int]int)
	var res1 []int
	var res2 []int

	for _, num := range nums1 {
		dic1[num] = 1
	}

	for _, num := range nums2 {
		dic2[num] = 1
	}

	for _, num := range nums1 {
		if _, ok := dic2[num]; !ok {
			if dic1[num] != 0 {
				res1 = append(res1, num)
				dic1[num] = 0
			}
		}
	}

	for _, num := range nums2 {
		if _, ok := dic1[num]; !ok {
			if dic2[num] != 0 {
				res2 = append(res2, num)
				dic2[num] = 0
			}
		}
	}

	return [][]int{res1, res2}
}