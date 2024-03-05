package main

import "fmt"

func main() {
	fmt.Println(minAddToMakeValid("())"))
	fmt.Println(minAddToMakeValid("((("))
}

func minAddToMakeValid(s string) int {
	res := 0
	// Qavzlarni ochilishlarini solish uchun list ochish
	var ls []string
	for i := 0; i < len(s); i++ {
		// Agar qavz oshilishi bo'lsa listga qo'shadi
		// Aks holda listdan qavz ochilishini olib tashlaydi
		if s[i] == '(' {
			ls = append(ls, string(s[i]))
		} else {
			// Agar list bo'sh bo'lsa, bu qavzni yopilishi ortiqcha degani
			// va uni hisoblash uchun res ga bir qo'shib ketadi
			if len(ls) > 0 {
				ls = ls[:len(ls)-1]
			} else {
				res++
			}
		}
	}
	// Yopilmagan qavzlarni listga qo'shadi
	res += len(ls)
	return res
}
