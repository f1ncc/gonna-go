package main

import (
	"fmt"
	"sync"

)


type Chopstick struct {
}


type Philosopher struct {
	id int
}

var pl = sync.Pool {
	New: func() interface{} {
		return new(Chopstick)
	},
}

var wg sync.WaitGroup

func (p Philosopher) eat(host chan int) {
	defer wg.Done()

	// to get permission from the host
	<-host

	fmt.Printf("starting to eat %d\n", p.id)

	left := pl.Get()
	right := pl.Get()

	pl.Put(left)
	pl.Put(right)
	fmt.Printf("finishing eating %d\n", p.id)

	host <- 1
}

func main() {

	host := make(chan int, 2)
	var NUM_EAT int = 3

	for i := 0; i < 5; i++ {
		pl.Put(new(Chopstick))
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < NUM_EAT; j++ {
			wg.Add(1)
			go philosophers[i].eat(host)
		}
	}

	host <- 1
	host <- 1

	wg.Wait()
	fmt.Printf("all of them have finished dining %d times concurrently", NUM_EAT)
}