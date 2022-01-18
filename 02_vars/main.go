package main

import "fmt"

func main() {

	// Using var
	// var name = "Jose Romario Peixoto"
	// var name string = "Jose Romario Peixoto"

	// name := "Jose Romario Peixoto"
	// email := "brad@gmail.com"

	name, email := "Jose Romario Peixoto", "romario@gmail.com"
	var age int32 = 23
	// var age int = 23
	// var age = 23

	const isCool = true
	// isCool = false

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
}
