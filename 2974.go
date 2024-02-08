package main

import "fmt"

func numberGame(nums []int) []int {
	var res []int
	arr := Sort_arr(nums)

	for i := 0; i < len(arr); i += 2 {
		res = append(res, arr[i+1])
		res = append(res, arr[i])
	}
	return res
}

func Sort_arr(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		ind := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[ind] {
				ind = j
			}
		}
		nums[i], nums[ind] = nums[ind], nums[i]
	}
	return nums
}

func main() {
	arr := []int{5, 4, 2, 3}
	fmt.Println(numberGame(arr))
}
