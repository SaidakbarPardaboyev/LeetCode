package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	res := float64(math.MinInt64)

	cursum := 0
	start := 0
	for end := 0; end < len(nums); end++ {
		cursum += nums[end]
		if end >= (k - 1) {
			res = math.Max(res, float64(cursum)/float64(k))
			cursum -= nums[start]
			start++
		}
	}
	return res
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))
}
