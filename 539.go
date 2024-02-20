package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(findMinDifference([]string{"12:12", "00:13"}))
	fmt.Println(findMinDifference([]string{"12:13", "00:12"}))
}

func findMinDifference(timePoints []string) int {
	res := math.MaxInt
	for i := 0; i < len(timePoints); i++ {
		for j := i + 1; j < len(timePoints); j++ {
			num1 := Get(timePoints[i])
			num2 := Get(timePoints[j])
			InerRes := num1 - num2
			if InerRes < 0 {
				InerRes = -InerRes
			}
			if InerRes < res {
				res = InerRes
			}
			fmt.Println("res: ", res)
		}
	}
	return res
}

func Get(timePoints string) int {
	var res int
	dic := map[string]int{
		"01": 1, "02": 2, "03": 3, "04": 4, "05": 5, "06": 6, "07": 7,
		"08": 8, "09": 9, "10": 10, "11": 11, "12": 12, "13": 13,
		"14": 14, "15": 15, "16": 16, "17": 17, "18": 18, "19": 19,
		"20": 20, "21": 21, "22": 22, "23": 23, "00": 24,
	}
	if timePoints[:2] == "00" {
		if timePoints[3:] == "00" {
			res = dic["00"] * 60
		} else {
			res, _ = strconv.Atoi(timePoints[3:])
		}
	} else {
		res = dic[timePoints[:2]] * 60
		tem, _ := strconv.Atoi(timePoints[3:])
		res += tem
	}
	return res
}
