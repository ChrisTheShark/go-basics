package concurrency

import (
	"log"
)

func main() {
	sum := puller(incrementor("Foo: "))
	sum2 := puller(incrementor("Bar: "))
	log.Printf("Final Counts: sum: %v, sum2: %v", <-sum, <-sum2)
}

func incrementor(name string) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
			log.Println(name, i)
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for i := range c {
			sum += i
		}
		out <- sum
		close(out)
	}()
	return out
}