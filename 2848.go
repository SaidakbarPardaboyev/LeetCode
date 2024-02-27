package main

import "fmt"

func main() {
	fmt.Println(numberOfPoints([][]int{{1, 9}, {2, 10}, {6, 7}, {8, 9}, {5, 8}, {1, 3}}))
	// fmt.Println(numberOfPoints([][]int{{1, 3}, {5, 8}}))
}

func numberOfPoints(nums [][]int) int {
	SortArr(nums)
	fmt.Println(nums)
	nums = Adding(nums)
	fmt.Println(nums)
	return Count(nums)
}

func Count(arr [][]int) int {
	sch := 0
	for i := 0; i < len(arr)-1; i++ {
		sch += arr[i+1][0] - arr[i][1] - 1
	}
	return arr[len(arr)-1][1] - arr[0][0] + 1 - sch
}

func SortArr(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j][0] < arr[j-1][0] {
			if arr[j][0] < arr[j-1][0] {
				tem := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = tem
			}
			j -= 1
		}
	}
}

func Adding(arr [][]int) [][]int {
	lamp := true
	for len(arr) > 1 {
		i := 0
		for i < len(arr)-1 {
			// [1, 5], [2, 8]
			if arr[i][1] > arr[i+1][0] && arr[i][1] <= arr[i+1][1] {
				fmt.Println("kirdi")
				arr[i][1] = arr[i+1][1]
				arr = append(arr[:i+1], arr[i+2:]...)
				// [1, 8], [2, 5]
			} else if arr[i][1] > arr[i+1][1] {
				arr = append(arr[:i+1], arr[i+2:]...)
				// [1, 5], [5, 8]
			} else if arr[i][1] == arr[i+1][0] {
				arr[i][1] = arr[i+1][1]
				arr = append(arr[:i+1], arr[i+2:]...)
			} else if arr[i][0] == arr[i+1][0] && arr[i][1] == arr[i+1][1] {
				arr = append(arr[:i+1], arr[i+2:]...)
			} else {
				lamp = false
			}
			i += 1
		}
		fmt.Println((arr))
		if !lamp {
			break
		}
	}
	return arr
}
