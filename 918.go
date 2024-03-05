package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2}))
}

func maxSubarraySumCircular(nums []int) int {
	Maxx := 0
	nums = append(nums, nums...)
	for i := 0; i < len(nums)/2; i++ {
		Sum_subarray := 0
		for j := i; j < len(nums)/2+i; j++ {
			Sum_subarray += nums[j]

			if Sum_subarray < 0 {
				Sum_subarray = 0
			}
			if Sum_subarray > Maxx {
				Maxx = Sum_subarray
			}
		}
	}
	if Maxx == 0 {
		return FindMaxx(&nums)
	}
	return Maxx
}

func FindMaxx(nums *[]int) int {
	maxx := (*nums)[0]
	for i := 0; i < len(*nums)/2; i++ {
		if (*nums)[i] > maxx {
			maxx = (*nums)[i]
		}
	}
	return maxx
}
