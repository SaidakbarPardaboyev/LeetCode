import "fmt"

func integerReplacement(n int) int {
	siklCounter := 0
	for n > 1 {
		fmt.Println(n)
		if n%2 == 0 {
			n /= 2
		} else {
			if (n+1)%4 == 0 && n != 3 {
				n += 1
			} else {
				n -= 1
			}
		}
		siklCounter++
	}
	return siklCounter
}