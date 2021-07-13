package main

import (
	"fmt"
	"time"
)

// SNIPPETSTART OMIT
type Ball struct{ hits int }

func main() {
	i := 0
	go player("player1", &i)
	go player("player2", &i)

	time.Sleep(2 * time.Second)
}

func player(name string, i *int) {
	for *i < 10 {
		*i++
		fmt.Println(name, *i)
		time.Sleep(100 * time.Millisecond)
	}
}

// SNIPPETEND OMIT
