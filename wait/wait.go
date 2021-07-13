package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	c := make(chan struct{}) // We don't need any data to be passed, so use an empty struct
	for i := 0; i < 100; i++ {
		go func() { // anonymous func
			doSomething()
			// signal that the routine has completed
			c <- struct{}{} // HLsync
		}()
	}

	// Since we spawned 100 routines, receive 100 messages.
	for i := 0; i < 100; i++ {
		<-c // HLsync
	}
	fmt.Println("DONE")
}

func doSomething() {
	time.Sleep(3 * time.Second)
}

// END OMIT
