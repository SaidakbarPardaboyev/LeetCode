package main

import (
	"fmt"
	"math"
)

func minimumSum(nums []int) int {
	res := -1
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] < nums[j] && nums[j] > nums[k] {
					if res == -1 {
						res = nums[i] + nums[j] + nums[k]
					} else {
						tem := nums[i] + nums[j] + nums[k]
						res = int(math.Min(float64(res), float64(tem)))
					}
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(minimumSum([]int{5, 4, 8, 7, 10, 2}))
}
