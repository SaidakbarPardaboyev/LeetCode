func mostFrequentPrime(mat [][]int) int {
	primeNumbersWithFrequency := map[int]int{}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			makedNumbers := makeNumbersByTraveling8Paths(mat, i, j)
			for _, number := range makedNumbers {
				if _, ok := primeNumbersWithFrequency[number]; ok || IsPrimeNumber(number) {
					primeNumbersWithFrequency[number]++
				}
			}
		}
	}

	numberOfMostFrequency := 0
	mostFrequentPrime := -1

	for number, frequency := range primeNumbersWithFrequency {
		if number > 10 && frequency > numberOfMostFrequency {
			numberOfMostFrequency = frequency
			mostFrequentPrime = number
		} else if number > 10 && frequency == numberOfMostFrequency {
			if number > mostFrequentPrime {
				numberOfMostFrequency = frequency
				mostFrequentPrime = number
			}
		}
	}

	return mostFrequentPrime
}

func makeNumbersByTraveling8Paths(nums [][]int, CurrentColumnIndex, CurrentRowIndex int) []int {
	directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}, {-1, 1}, {-1, -1}, {1, 1}, {1, -1}}
	numbers := []int{}

	for _, direction := range directions {
		directionForI := direction[0]
		directionForJ := direction[1]

		numberMakedTraveling := 0

		i := CurrentColumnIndex
		j := CurrentRowIndex
		for i >= 0 && i < len(nums) && j >= 0 && j < len(nums[CurrentColumnIndex]) {
			numberMakedTraveling = numberMakedTraveling*10 + nums[i][j]
			i += directionForI
			j += directionForJ
			numbers = append(numbers, numberMakedTraveling)
		}
	}

	return numbers
}

func IsPrimeNumber(number int) bool {
	if number == 1 {
		return false
	}

	if number == 2 {
		return true
	}

	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}