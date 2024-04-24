import "math"

func maxSubArray(nums []int) int {
	preNumber := 0
	maxSum := math.MinInt64

	for _, number := range nums {
		if preNumber+number > number {
			preNumber += number
		} else {
			preNumber = number
		}
		maxSum = max(maxSum, preNumber)
	}

	return maxSum
}