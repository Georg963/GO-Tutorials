package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var wg sync.WaitGroup
var mut sync.Mutex
var i = 0

func main() {
	fmt.Println(i)
	wg.Add(3)
	go inc()
	go inc()
	go inc()
	wg.Wait()
	fmt.Println(i)
}

func inc() {
	once.Do(func() { fmt.Println("GO!") })
	mut.Lock()
	i++
	mut.Unlock()
	wg.Done()
}
