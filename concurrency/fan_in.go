package concurrency

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	in := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-in)
	}
	fmt.Println("You guys are boring, I'm leaving.")
}

func boring(msg string) <-chan string {
	out := make(chan string)
	go func() {
		for i := 0; ; i++ {
			out <- fmt.Sprintf("%v: %v", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return out
}

func fanIn(input1, input2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for {
			out <- <-input1
		}
	}()
	go func() {
		for {
			out <- <-input2
		}
	}()
	return out
}