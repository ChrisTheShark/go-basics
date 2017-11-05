package basics

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName string
	Age int
}

func (person Person) getFullName() string {
	return person.FirstName + person.LastName
}
func (person Person) String() string {
	return fmt.Sprintf("[%v %v, %v]", person.FirstName, person.LastName, person.Age)
}

type DoubleZero struct {
	Person
	FirstName string
	LicenseToKill bool
}

func (doubleZero DoubleZero) getFullName() string {
	return doubleZero.FirstName
}

func main() {
	spy := DoubleZero{
		Person{"Json", "Johns", 32},
		"007",
		true,
	}
	fmt.Println(spy.FirstName, spy.Person.FirstName)
	fmt.Println(spy.getFullName())

	var person Person
	fmt.Println(person)
	fmt.Println(spy.Person)

	fmt.Println(spy.(Person))
}
