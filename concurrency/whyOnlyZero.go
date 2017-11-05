package concurrency

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	// Simple fmt.Println below only takes the first value from the
	// channel and exits. To get all values we must loop and read more
	// values.
	//fmt.Println(<-c)
	for i := range c {
		fmt.Println(i)
	}
}