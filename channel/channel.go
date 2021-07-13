package main

import (
	"time"
	"fmt"
)

// START OMIT
func doSomething(done chan bool) {
	time.Sleep(3*time.Second)
	fmt.Println("did something")
	done <- true
}

func main() {
	done := make(chan bool)
	go doSomething(done)
	fmt.Println("work somewhere else")
	<-done
}
// END OMIT

