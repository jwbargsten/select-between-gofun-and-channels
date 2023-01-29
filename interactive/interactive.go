package main

import (
	"time"
	"fmt"
)

func walk(message string) {
	time.Sleep(3*time.Second)
	fmt.Println(message)
}

func main() {

	go walk("walked through the park")
	go walk("walked through Amsterdam")
	time.Sleep(4*time.Second)
}
