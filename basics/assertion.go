package basics

import "fmt"

type vehicle interface {}

type car struct {
	vehicle
	NumDoors int
}

type boat struct {
	vehicle
	Length float64
}


func main() {
	var name interface{} = "Sydney"
	str, ok := name.(string)
	fmt.Println(str, ok)

	getInfo(car{NumDoors: 3})
	getInfo(boat{Length: 14.4})
}

func getInfo(v vehicle) {
	if c, ok := v.(car); ok {
		fmt.Println(c.NumDoors)
	} else if b, ok := v.(boat); ok {
		fmt.Println(b.Length)
	} else {
		fmt.Println("Not a boat or a car.")
	}
}