package basics

import "fmt"

type Animal interface {
	makeNoise() string
}

type Dog struct {
	Bark string
}

func (dog Dog) makeNoise() string {
	return dog.Bark
}

func main()  {
	dog := Dog{"Woof!"}
	fmt.Println(dog.makeNoise())
	approachAnimal(dog)
}

func approachAnimal(animal Animal) {
	fmt.Println(animal.makeNoise())
}