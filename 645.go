package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findErrorNums([]int{1, 5, 3, 2, 2, 7, 6, 4, 8, 9}))
}

func findErrorNums(nums []int) []int {
	sort.Ints(nums)
	num := 1
	lost := -1
	fmt.Println(nums)
	for i := 0; i < len(nums); {
		if nums[i] == num {
			num++
		} else {
			if i < len(nums)-1 {
				i++
				if nums[i] == num {
					num++
				} else {
					lost = num
					break
				}
			}
		}
		i++
	}
	num = 0
	for i := range nums {
		if nums[i] == num {
			break
		}
		num = nums[i]
	}
	if lost == -1 {
		lost = nums[len(nums)-1] + 1
	}
	return []int{num, lost}
}
