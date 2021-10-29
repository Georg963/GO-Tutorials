package main

import "fmt"

func main() {
	//1
	k := 0
	for k < 3 {
		fmt.Println(k)
		k++
	}
	fmt.Println()

	//2
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	//3
	var student = map[string]int{"George": 26, "Stephan": 24}
	for key, value := range student {
		fmt.Printf("Key: -> %v, Value: -> %v\n", key, value)
	}

	//4
	var slice = []string{"George", "Dave", "Christian"}
	for _, val := range slice {
		fmt.Println(val)
	}
}
