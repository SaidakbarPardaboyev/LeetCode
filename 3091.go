import "math"

func minOperations(k int) int {
	number := 1
	numberOfOperations := math.MaxInt64

	if number >= k {
		return 0
	}

	CurrentNumberOfOperations := 0
	for number <= k {
		remainder := 0
		if k%number > 0 {
			add++
		}
		numberOfOperations = min(numberOfOperations, CurrentNumberOfOperations+k/number+add)
		CurrentNumberOfOperations++
		number++
	}
	return numberOfOperations - 1
}