package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
}

func findRestaurant(list1 []string, list2 []string) []string {
	resSum := math.MaxInt
	var res []string
	for i := 0; i < len(list1); i++ {
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				if resSum > i+j {
					res = []string{list1[i]}
					resSum = i + j
				} else if resSum == i+j {
					res = append(res, list1[i])
				}
			}
		}
	}
	return res
}
