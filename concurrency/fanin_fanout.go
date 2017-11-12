package concurrency

import (
	"fmt"
	"sync"
)

func main() {
	start := generate(2, 3)

	for n := range merge(square(start), square(start)) {
		fmt.Println(n)
	}
}

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, i := range nums {
			out <- i
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out
}

func merge(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))

	for _, c := range ins {
		go func(ch <-chan int) {
			for i := range ch {
				out <- i
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}