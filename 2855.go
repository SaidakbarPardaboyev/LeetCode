package main

import "fmt"

func main() {
	fmt.Println(minimumRightShifts([]int{3, 4, 5, 1, 2}))
}

func minimumRightShifts(nums []int) int {
	sch := 0
	if Check(nums) {
		return sch
	}
	for i := 0; i < len(nums); i++ {
		sch += 1
		nums = roteArr(nums)
		if Check(nums) {
			return sch
		}
	}
	return -1
}

func roteArr(nums []int) []int {
	arr := make([]int, len(nums))
	arr[0] = nums[len(nums)-1]

	for i := 0; i < len(nums)-1; i++ {
		arr[i+1] = nums[i]
	}
	return arr
}

func Check(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if !(nums[i] < nums[i+1]) {
			return false
		}
	}
	return true
}
