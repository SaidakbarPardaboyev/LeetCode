package main

import (
	"fmt"
	"strconv"
)

func solveEquation(equation string) string {
	var ls1 [][]int
	var ls2 [][]int
	for i := 1; i < len(equation); i++ {
		if equation[i] == '=' {
			ls1 = SplitString(equation[:i+1])
			ls2 = SplitString(equation[i+1:] + string('='))
			break
		}
	}
	val := 0
	x := 0
	for i := range ls1 {
		if ls1[i][0] == 0 {
			x += ls1[i][1]
		} else {
			val -= ls1[i][1]
		}
	}

	for i := range ls2 {
		if ls2[i][0] == 0 {
			x -= ls2[i][1]
		} else {
			val += ls2[i][1]
		}
	}
	if x == 0 {
		return "Infinite solutions"
	} else {
		if val%x == 0 {
			st := "x="
			tem := strconv.Itoa(val / (-x))
			return st + string(tem)
		}
	}
	return ""
}

func SplitString(st string) [][]int {
	var res [][]int
	tem := ""
	for i := range st {
		if st[i] == '+' || st[i] == '-' || st[i] == '=' {
			if len(tem) > 0 {
				if tem[len(tem)-1] == 'x' {
					soni, _ := strconv.Atoi(tem[:len(tem)-1])
					res = append(res, []int{0, soni})
				} else {
					soni, _ := strconv.Atoi(tem)
					res = append(res, []int{1, soni})
				}
			}
			tem = string(st[i])
		} else if st[i] == 'x' {
			if len(tem) == 0 || len(tem) == 1 {
				tem += "1"
			}
			tem += string(st[i])
		} else {
			tem += string(st[i])
		}
	}
	return res
}

func main() {
	fmt.Println(solveEquation("2x+3x-6x=x+2"))
}
