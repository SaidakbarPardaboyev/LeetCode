package main

import "fmt"

func main() {
	fmt.Println(hasAlternatingBits(2147483648))
}

func hasAlternatingBits(n int) bool {
	var Bdig []int

	for n > 0 {
		Bdig = append(Bdig, n%2)
		n /= 2
	}

	for i := 0; i < len(Bdig)-1; i++ {
		if Bdig[i] == Bdig[i+1] {
			return false
		}
	}

	return true
}
