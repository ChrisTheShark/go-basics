package main

import "fmt"

func main() {
	fmt.Println(<-calculateFactorial(4))
}

func calculateFactorial(num int) chan int {
	out := make(chan int)
	go func() {
		out <- factorial(num)
		close(out)
	}()
	return out
}

func factorial(num int) int {
	if num != 0 {
		return num * factorial(num - 1)
	}
	return 1
}