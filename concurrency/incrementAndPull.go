package concurrency

import "fmt"

func main() {
	c := incrementor()
	for i := range reducer(c) {
		fmt.Println(i)
	}
}

func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func reducer(pipe <-chan int) <-chan int {
	out := make(chan int)
	var sum int
	go func() {
		for i := range pipe {
			sum += i
		}
		out <- sum
		close(out)
	}()
	return out
}