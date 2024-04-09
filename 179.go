func largestNumber(nums []int) string {
	numsString := []string{}

	for _, num := range nums {
		numsString = append(numsString, strconv.Itoa(num))
	}

	res := strings.Join(QuickSort(numsString), "")
	caseZero, _ := strconv.Atoi(res)
	if caseZero == 0 {
		return "0"
	}
	return res
}

func QuickSort(nums []string) []string {
	if len(nums) < 2 {
		return nums
	}
	ind := len(nums) / 2
	pivot := nums[ind]
	less := []string{}
	greater := []string{}

	for indx, num := range nums {
		if indx == ind {
			continue
		}
		checked := CheckGreater(pivot, num)
		if checked == 0 || checked == 2 {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}

	less = QuickSort(less)
	greater = QuickSort(greater)
	greater = append(greater, pivot)
	greater = append(greater, less...)
	return greater
}

func CheckGreater(pivot, num string) int {
	ab := pivot + num
	ba := num + pivot

	if ab > ba {
		return 0
	} else if ba > ab {
		return 1
	} else {
		return 2
	}
}