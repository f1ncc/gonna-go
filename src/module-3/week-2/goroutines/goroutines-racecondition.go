package main

import (
	"fmt"
	"time"
)

func f1() {
	fmt.Println("sup")
}

func f2() {
	fmt.Println("mate")
}


func main() {
	
	for i:=0; i<10; i++ {
		go f1()
		go f2()
	}

	time.Sleep(100 * time.Millisecond)
}

// adding delay is one of the simplest ways to examine the race condition of these two goroutines. when this program is executed multiple times, observe the pattern in which
// they print the strings supplied as args. it is almost random. this is because of the race condition and the interleavings between the two goroutines. everytime before you execute this program
// the race conditon at the time, is in a sequence which we can get to know by executing the program and looking at the pattern in which the strings are being printed.