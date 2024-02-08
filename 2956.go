package main

import "fmt"

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	var res []int
	sch := 0
	for _, i := range nums1 {
		for _, j := range nums2 {
			if i == j {
				sch++
				break
			}
		}
	}
	res = append(res, sch)

	sch = 0
	for _, i := range nums2 {
		for _, j := range nums1 {
			if i == j {
				sch++
				break
			}
		}
	}
	res = append(res, sch)

	return res
}

func main() {
	num1 := []int{4, 3, 2, 3, 1}
	num2 := []int{2, 2, 5, 2, 3, 6}
	fmt.Println(findIntersectionValues(num1, num2))
}
