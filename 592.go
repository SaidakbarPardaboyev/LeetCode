package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fractionAddition(("-1/2+1/2")))
}

func fractionAddition(expression string) string {
	ls := MakeArif(expression + "+")
	res := ls[0]
	for i := 1; i < len(ls); i++ {
		res = Calculate(res, ls[i])
	}
	for i := 2; i < res[1]+1; i++ {
		for res[0]%i == 0 && res[1]%i == 0 {
			if res[0]%i == 0 && res[1]%i == 0 {
				res[0] /= i
				res[1] /= i
			}
		}
	}
	numerator := strconv.Itoa(res[0])
	denumerator := strconv.Itoa(res[1])
	if numerator == "0" && denumerator != "0" {
		denumerator = "1"
	} else if numerator != "0" && denumerator == "0" {
		denumerator = "1"
	}
	return string(numerator) + "/" + string(denumerator)
}

func Calculate(num1 []int, num2 []int) []int {
	if num1[1] != num2[1] {
		ind1 := num1[1]
		ind2 := num2[1]

		num1[0] *= ind2
		num1[1] *= ind2

		num2[0] *= ind1
		num2[1] *= ind1
	}

	res := []int{num1[0] + num2[0], num1[1]}

	return res
}

func MakeArif(exp string) [][]int {
	var ls [][]int
	sNum := ""
	var num []int
	for i := 0; i < len(exp); i++ {
		if exp[i] == '/' {
			tem, _ := strconv.Atoi(sNum)
			num = append(num, tem)
			sNum = ""
		} else if (exp[i] == '+' || exp[i] == '-') && i != 0 {
			tem, _ := strconv.Atoi(sNum)
			num = append(num, tem)
			ls = append(ls, num)
			num = []int{}
			sNum = string(exp[i])
		} else {
			sNum += string(exp[i])
		}
	}
	return ls
}
