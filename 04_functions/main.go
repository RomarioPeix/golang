package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	fmt.Println(greeting("Romário "))
	fmt.Println(getSum(3, 4))
}
