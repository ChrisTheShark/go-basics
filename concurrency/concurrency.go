package concurrency

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

var wg sync.WaitGroup

func main() {
	// Add two items to the wait group.
	wg.Add(2)
	go foo()
	go bar()
	// Wait for above functions to report done.
	wg.Wait()

	fmt.Println(runtime.NumCPU())
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("FOO: ", i)
		// Sleep three milliseconds to ensure some interleaving.
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	// Report done.
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("BAR: ", i)
		// Sleep three milliseconds to ensure some interleaving.
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	// Report done.
	wg.Done()
}