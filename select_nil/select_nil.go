package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	c1 := make(chan int)
	c2 := make(chan int)
	go produce(c1)
	go produce(c2)

	var c2Current chan int
	for i:= 0; i<20; i++{
		select {
		case i1 := <-c1:
			fmt.Println("reading c1:", i1)
			if i1 > 5 && c2Current == nil {
				fmt.Println("switching c2 on")
				c2Current = c2
			}
		case i2 := <-c2Current:
			fmt.Println("reading c2:", i2)

		}
	}
	// END OMIT
}

func produce(c chan int) {
	i := 0
	for ; i < 10; i++ {
		c <- i
		time.Sleep(100 * time.Millisecond)
	}
}
