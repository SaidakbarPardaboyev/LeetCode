package main

import (
	"fmt"
)

func main() {
	fmt.Println(countPrimeSetBits(6, 10))
	fmt.Println(countPrimeSetBits(6, 100000))
}

func countPrimeSetBits(left int, right int) int {
	res := 0
	for i := left; i < right+1; i++ {
		res += FindBin(i)
	}
	return res
}

func FindBin(num int) int {
	res := 0
	for num > 0 {
		tem := num % 2
		if tem == 1 {
			res += 1
		}
		num /= 2
	}
	if IsPrimary(res) {
		return 1
	}
	return 0
}

func IsPrimary(num int) bool {
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
