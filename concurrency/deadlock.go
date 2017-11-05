package concurrency

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 3
	}()
	// Results in deadlock without the write in a go routine. This causes
	// go to launch the routine and wait below for the write.
	fmt.Println(<-c)
}