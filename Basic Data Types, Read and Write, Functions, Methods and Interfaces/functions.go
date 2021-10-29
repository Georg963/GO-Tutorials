package main

import (
	"errors"
	"fmt"
)

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("ZeroDivisionError")
	}
	return x / y, nil
}

type add func(a, b int) int

func funcAsArgument(f func(a, b int) int, c int) int {
	return f(2, 3) + c
}

func returningFunc() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func main() {
	//test 1
	fmt.Println(div(2, 0))
	//test 2
	fmt.Println(div(4, 2))

	//first class functions
	func() {
		fmt.Println("Hello World!")
	}()

	a := func() int {
		return 5 + 3
	}()
	fmt.Println(a)

	b := func(a, b int) int {
		return a * b
	}(2, 3)
	fmt.Println(b)

	//User defined function types
	var c add = func(a, b int) int { return a + b }
	fmt.Println(c(1, 2))

	//Higher-order functions
	//1. Passing functions as arguments to other functions
	d := func(a, b int) int {
		return a + b
	}
	fmt.Println(funcAsArgument(d, 3))

	//2. Returning functions from other functions
	e := returningFunc()
	fmt.Println(e(2, 3))

	//Closures
	g := 5
	func() {
		fmt.Println(g + 1)
	}()
}
