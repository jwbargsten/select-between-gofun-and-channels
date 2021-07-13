package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1)
			doSomething(&wg)
		}()
	}
	wg.Wait()
	fmt.Println("DONE")
}

func doSomething(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
}

// END OMIT
