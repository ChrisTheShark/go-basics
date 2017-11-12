package main

import (
	"fmt"
)

func main() {
	for i := range factorialLoop(generate(100)) {
		fmt.Println(i)
	}
}

func generate(num int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < num; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorialLoop(in <-chan int) <-chan float64 {
	out := make(chan float64)
	go func() {
		for i := range in {
			factorialValue := 1.0
			for j := i; j > 0; j-- {
				factorialValue *= float64(j)
			}
			out <- factorialValue
		}
		close(out)
	}()
	return out
}