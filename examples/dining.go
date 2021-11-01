package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type host struct {
	eaters chan bool
}

type philosopher struct {
	id int
}

type chopstick struct {
}

var chopsticks = make(chan chopstick, 5)
var wg sync.WaitGroup

func (p philosopher) eat(h *host) {
	for i := 0; i < 3; i++ {
		<-h.eaters

		left, right := <-chopsticks, <-chopsticks

		fmt.Println("starting to eat: ", p.id)

		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		fmt.Println("finished eating: ", p.id)

		chopsticks <- right
		chopsticks <- left

		h.eaters <- true
	}
	wg.Done()
}

func (h host) begin() {
	for i := 0; i < 2; i++ {
		h.eaters <- true
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	host := new(host)
	host.eaters = make(chan bool, 2)

	go host.begin()

	philosophers := make([]*philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &philosopher{i + 1}
		chopsticks <- *new(chopstick)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosophers[i].eat(host)
	}

	wg.Wait()
}
