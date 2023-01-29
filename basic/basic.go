package main

import (
	"fmt"
	"time"
)

func walk(notify chan int, msg string) {
	time.Sleep(3 * time.Second)
	fmt.Println(msg)
	notify <- 1
}

func main() {
	notify := make(chan int)

	go walk(notify, "done1")
	go walk(notify, "done2")

	counts := 0
	for {
		if counts == 2 {
			break
		}

		select {
		case <-notify:
			counts++
		}
	}
}
