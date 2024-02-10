package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findIndices([]int{5, 1, 4, 1}, 2, 4))
}

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	res := []int{-1, -1}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if int(math.Abs(float64(i)-float64(j))) >= indexDifference && int(math.Abs(float64(nums[i])-float64(nums[j]))) >= valueDifference {
				res[0] = i
				res[1] = j
				return res
			}
		}
	}
	return res
}
