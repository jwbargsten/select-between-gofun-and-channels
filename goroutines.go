package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
	println("DONE")
}

// blocking
// when does a channel block
// reading/writing
// flow/for loop
// https://tour.golang.org/concurrency/11
// goroutine -> & in shell (bg) (src: https://www.youtube.com/watch?v=f6kdp27TYZs, 7:41)
// Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives. (See the next slide.) 

//https://talks.golang.org/2013/advconc.slide#1

// cannels convey data, timer events, cancellation signals

// Goroutines serialise access to local mutable state
// stack traces & deadlock detector
// race detector
