package main

import "fmt"

func main() {
	x := 5
	if x == 4 {
		fmt.Printf("x is %v\n", x)
	} else if x == 5 {
		fmt.Printf("x is %v\n", x)
	} else {
		fmt.Println("x is neither 4 nor 5")
	}
}
