package basics

import "fmt"

func main() {
	inspect(3)
	inspect("Charlie")
	inspect(46.66)
	inspect([]string{"Jon"})
	inspect(map[string]int{"jon": 23})
}

func inspect(arg interface{}) {
	fmt.Printf("%T\n", arg)
}