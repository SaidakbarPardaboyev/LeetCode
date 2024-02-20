package main

import (
	"fmt"
)

func main() {
	fmt.Println(arrayPairSum([]int{1, 4, 3, 2}))
}

func arrayPairSum(nums []int) int {
	// slices.Sort(nums)
	Sort(nums)
	sum := nums[0]
	for i := 2; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func Sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		ind := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[ind] {
				ind = j
			}
		}
		nums[ind], nums[i] = nums[i], nums[ind]
	}
}
