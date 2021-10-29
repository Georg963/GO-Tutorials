package main

import "fmt"

func inc1(x int) {
	x++
}

func inc2(x *int) {
	*x++
}

func main() {
	i := 1
	inc1(i)
	fmt.Println(i)
	fmt.Println("-----")
	inc2(&i)
	fmt.Println(i)

	fmt.Println("-----")
	f := 3
	n := &f
	fmt.Println(*n)
	*n = 2
	fmt.Println(f)
}
