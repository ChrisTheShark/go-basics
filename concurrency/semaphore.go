package concurrency

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()
	go func() {
		<-done
		<-done
		close(c)
		close(done)
	}()

	for i := range c {
		fmt.Println(i)
	}
}