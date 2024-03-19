package main

import (
	"fmt"
	"sort"
)

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)

	res := 0
	end := len(nums) - 1
	for start := 0; start < end; {
		if nums[start]+nums[end] == k {
			res++
			start++
			end--
		} else if nums[start]+nums[end] < k {
			start++
		} else {
			end--
		}
	}

	return res
}

func main() {
	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6))
}
