package main

import "fmt"

func findPeaks(mountain []int) []int {
	var ls []int
	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			ls = append(ls, i)
		}
	}
	return ls
}

func main() {
	arr := []int{1, 4, 3, 8, 5}
	fmt.Println(findPeaks(arr))
}
