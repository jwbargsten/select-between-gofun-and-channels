package main

import (
	"time"
	"fmt"
)

// START OMIT
func doSomething() {
	time.Sleep(3*time.Second)
	fmt.Println("did something")
}

func main() {
	go doSomething()
	fmt.Println("work somewhere else") 
	time.Sleep(4*time.Second)
}
// END OMIT
