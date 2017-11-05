package concurrency

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var wg sync.WaitGroup
var counter int = 0

func main() {
	wg.Add(2)
	go increment("FOO: ")
	go increment( "BAR: ")
	wg.Wait()
	fmt.Printf("FINAL COUNT: %v", counter)
}

func increment(name string) {
	for i := 0; i < 45; i++ {
		mutex.Lock()
		counter++
		fmt.Printf("%v updates counter to %v\n", name, counter)
		mutex.Unlock()
	}
	wg.Done()
}