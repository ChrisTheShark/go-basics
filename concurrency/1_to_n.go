package concurrency

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)
	n := 10

	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func() {
			for i := range c {
				fmt.Println(i)
			}
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}
}
