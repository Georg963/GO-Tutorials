package main

import "fmt"

type MyInt int

func (x MyInt) add(y MyInt) MyInt {
	return x + y
}

func main() {
	a := MyInt(3)
	b := MyInt(4)
	fmt.Println(a.add(b))
}
