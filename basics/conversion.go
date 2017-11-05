package basics

import "fmt"

func main()  {
	bs := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(bs)
	fmt.Println(string(bs))

	fmt.Println([]byte("hello"))
}