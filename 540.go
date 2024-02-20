package main

import "fmt"

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
}
func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	for i, j := 0, len(nums)-1; i <= j; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
		if nums[j] != nums[j-1] {
			return nums[j]
		}
		j -= 2
	}
	return 0
}
