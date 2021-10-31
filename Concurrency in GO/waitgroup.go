package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	str := "Hello World"

	wg.Add(2)
	go printFunc(str)
	go printFunc(str)
	wg.Wait()
}

func printFunc(str string) {
	fmt.Println(str)
	wg.Done()
}
