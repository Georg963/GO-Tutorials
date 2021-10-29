package main

import "fmt"

func sum(a ...int) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}
