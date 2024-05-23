package main

import "fmt"

func main() {
	ch := make(chan int, 0)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}
