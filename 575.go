package main

import "fmt"

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
}

func distributeCandies(candyType []int) int {
	ls := MakeSet(candyType)
	if len(ls) >= len(candyType)/2 {
		return len(candyType) / 2
	} else if len(ls) < len(candyType)/2 {
		return len(ls)
	} else {
		return 0
	}
}

func MakeSet(nums []int) []int {
	dic := map[int]int{}

	for i := range nums {
		dic[nums[i]] = nums[i]
	}
	var ls []int
	for i := range dic {
		ls = append(ls, dic[i])
	}
	fmt.Println(ls)
	return ls
}
