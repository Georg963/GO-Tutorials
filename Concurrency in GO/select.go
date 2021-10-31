package main

import "fmt"

func main() {
	chan_1 := make(chan int)
	chan_2 := make(chan int)

	go sum(chan_1, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	go sum(chan_2, 1, 2, 3, 4)

	select {
	case a := <-chan_1:
		fmt.Println(a)
	case b := <-chan_2:
		fmt.Println(b)
	}
}

func sum(c chan int, s ...int) {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	c <- sum
}
