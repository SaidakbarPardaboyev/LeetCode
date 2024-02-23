package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumSwap(2176))
}

func maximumSwap(num int) int {
	res := num
	var nums []int
	max := FindMax(num, &nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] < max {
			for j := i + 1; j < len(nums); j++ {
				tem := MakeSingleNum(nums, i, j)
				if tem > res {
					res = tem
				}
			}
		}
	}
	return res
}

func MakeSingleNum(nums []int, ind1 int, ind2 int) int {
	res := 0
	for i := range nums {
		if i != ind1 && i != ind2 {
			res = res*10 + nums[i]
		} else {
			if i == ind1 {
				res = res*10 + nums[ind2]
			} else if i == ind2 {
				res = res*10 + nums[ind1]
			}
		}
	}
	return res
}

func FindMax(num int, nums *[]int) int {
	max := num % 10
	*nums = append(*nums, num%10)
	num /= 10

	for num > 0 {
		if num%10 > max {
			max = num % 10
		}
		*nums = append(*nums, num%10)
		num /= 10
	}
	j := len(*nums) - 1
	for i := 0; i < j; i++ {
		(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
		j--
	}
	return max
}
