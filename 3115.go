func maximumPrimeDifference(nums []int) int {
	indOfFirstPrime := -1
	indOfLastPrime := -1

	i := 0
	for i < len(nums) {
		if IsPrimeNumber(nums[i]) {
			indOfFirstPrime = i
			break
		}
		i++
	}
	j := len(nums) - 1
	for j >= 0 {
		if IsPrimeNumber(nums[j]) {
			indOfLastPrime = j
			break
		}
		j--
	}
	return indOfLastPrime - indOfFirstPrime
}

func IsPrimeNumber(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}