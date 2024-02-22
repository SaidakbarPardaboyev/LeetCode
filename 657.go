package main

import "fmt"

func main() {
	fmt.Println(judgeCircle("UD"))
}

func judgeCircle(moves string) bool {
	i := 0
	j := 0
	for k := 0; k < len(moves); k++ {
		if moves[k] == 'U' {
			i++
		} else if moves[k] == 'D' {
			i--
		} else if moves[k] == 'R' {
			j++
		} else if moves[k] == 'L' {
			j--
		}
	}
	return i == 0 && j == 0
}
