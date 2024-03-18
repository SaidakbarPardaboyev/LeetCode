package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	small := math.MaxInt64
	medium := small

	for i := range nums {
		if nums[i] <= small {
			small = nums[i]
		} else if nums[i] <= medium {
			medium = nums[i]
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{1, 5, 0, 4, 1, 3}))
}
