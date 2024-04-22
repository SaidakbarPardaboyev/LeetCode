import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})

	maximumHappiness := 0
	happinessDecreasingBy := 0
	for i := 0; i < len(happiness) && k > 0; i++ {
		if happiness[i]-happinessDecreasingBy >= 0 {
			maximumHappiness += happiness[i] - happinessDecreasingBy
		}

		happinessDecreasingBy++
		k--
	}

	return int64(maximumHappiness)
}