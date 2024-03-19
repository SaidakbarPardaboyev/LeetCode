package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	maxWater := math.MinInt64
	i := 0
	j := len(height) - 1
	for i < j {
		tem := int(math.Min(float64(height[i]), float64(height[j]))) * (j - i)
		if tem > maxWater {
			maxWater = tem
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxWater
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
