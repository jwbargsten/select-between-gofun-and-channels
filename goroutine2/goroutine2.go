package main

import (
	"fmt"
	"time"
)

// START OMIT
func walk(msg string) {
	time.Sleep(3 * time.Second)
	fmt.Println(msg)
}

func main() {
	go walk("What a beautiful walk!")
	go walk("I had rain, not going to walk again, ever!")

	time.Sleep(4 * time.Second)
}
// END OMIT
