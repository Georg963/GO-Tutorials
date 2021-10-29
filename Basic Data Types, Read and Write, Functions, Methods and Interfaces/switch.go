package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	switch x {
	case 0:
		fmt.Println("x is 0")
	case 1:
		fmt.Println("x is 1")
	default:
		fmt.Println("x is neither 0 nor 1")
	}

	fmt.Println("-----")

	/*fallthrough
	keyword used to force the execution flow to fall
	through the successive case block.
	*/

	switch x {
	case 0:
		fmt.Println("x is 0")
	case 1:
		fmt.Println("x is 1")
		fallthrough
	case 2:
		fmt.Println("x is 2")
	case 3:
		fmt.Println("x is 3")
	default:
		fmt.Println("x is neither 0 nor 1")
	}
}
