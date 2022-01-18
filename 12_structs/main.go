package main

import (
	"fmt"
	"strconv"
)

// define person struct

type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

// greeting method (value receiver)
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age)
}

// hasBirthday method (pointer receiver)
func (person *Person) hasBirthday() {
	person.age++
}

// getMarried (pointer receiver)
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "M" {
		return
	} else {
		person.lastName = spouseLastName
	}
}

func main() {
	// init person using struct
	person1 := Person{
		firstName: "Jose Romario",
		lastName:  "Peixoto",
		city:      "Praia Grande",
		gender:    "M",
		age:       23,
	}

	person2 := Person{
		firstName: "Eloa Perici",
		lastName:  "Campos",
		city:      "Rio de Janiero",
		gender:    "F",
		age:       30,
	}
	// person1 := Person{
	// 	"Jose Romario",
	// 	"Peixoto",
	// 	"Praia Grande",
	// 	"M",
	// 	23,
	// }
	// fmt.Println(person1)

	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	// fmt.Println(person1.greet())
	// person1.hasBirthday()
	// person1.hasBirthday()
	// person1.hasBirthday()
	// person1.hasBirthday()
	// fmt.Println(person1.greet())

	fmt.Println(person1)
	person1.getMarried("Campos")
	fmt.Println(person1)

	fmt.Println(person2)
	person2.getMarried("Peixoto")
	fmt.Println(person2)

}
