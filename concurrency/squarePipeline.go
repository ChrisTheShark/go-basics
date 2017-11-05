package concurrency

import "fmt"

func main() {
	output(square(generate(2, 4, 6)))
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

func output(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
