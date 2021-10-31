package main

import "fmt"

func main() {
	c := make(chan int)
	c_2 := make(chan int, 2)
	go sum(c, 1, 2, 3, 4, 5, 6, 7)
	sums := <-c
	fmt.Println(sums)

	go sum(c_2, 1, 2, 3, 4, 5, 6, 7)
	go sum(c_2, 1, 2, 3, 4, 5, 6, 7)
	a, b := <-c_2, <-c_2
	fmt.Println(a, "----", b)
}

func sum(c chan int, s ...int) {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	c <- sum
}
