package main

import "fmt"

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	fmt.Println(findLHS([]int{1, 2, 3, 4}))
	fmt.Println(findLHS([]int{1, 1, 1, 1}))
}

func findLHS(nums []int) int {
	dic := GroupBy(nums)
	ls := Sort(dic)
	res := 0
	for i := 1; i < len(ls); i++ {
		if ls[i][0]-ls[i-1][0] == 1 {
			if ls[i][1]+ls[i-1][1] > res {
				res = ls[i][1] + ls[i-1][1]
			}
		}
	}
	return res
}

func Sort(dic map[int]int) [][]int {
	var ls [][]int
	for key, val := range dic {
		ls = append(ls, []int{key, val})
	}

	for i := 0; i < len(ls); i++ {
		ind := i
		for j := i + 1; j < len(ls); j++ {
			if ls[ind][0] > ls[j][0] {
				ind = j
			}
		}
		ls[ind], ls[i] = ls[i], ls[ind]
	}
	return ls
}

func GroupBy(ls []int) map[int]int {
	dic := map[int]int{}
	for i := range ls {
		if Check(ls[i], dic) == true {
			dic[ls[i]] += 1
		} else {
			dic[ls[i]] = 1
		}
	}
	return dic
}

func Check(key int, dic map[int]int) bool {
	for i := range dic {
		if i == key {
			return true
		}
	}
	return false
}
