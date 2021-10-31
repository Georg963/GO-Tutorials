package main

import (
	"fmt"
	"time"
)

func main() {
	str := "Hello World"
	go printFunc(str)
	time.Sleep(time.Millisecond * 50)
}

func printFunc(str string) {
	fmt.Println(str)
}
