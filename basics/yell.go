package basics

import (
	"fmt"
	"strings"
)

type Yell string

func (yell Yell) shout() string {
	return strings.ToUpper(string(yell))
}

func main() {
	var scream Yell
	scream = "hello"

	fmt.Printf("%T - %v\n", scream, scream.shout())
}