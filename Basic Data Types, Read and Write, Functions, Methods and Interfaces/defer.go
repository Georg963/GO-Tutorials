package main

import "fmt"

func main() {
	defer fmt.Println("deffered Functions calls!")
	fmt.Println("Hello World!")

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}
