package main

import "fmt"

func main() {
	ids := []int{33, 11, 14, 16, 13, 1}

	// loop through ids
	for i, id := range ids {
		fmt.Printf("number %d id %d\n", i, id)
	}

	// not using index
	for _, id := range ids {
		fmt.Printf("id %d\n", id)
	}

	// add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// range with map
	emails := map[string]string{
		"roms":    "roms@gmail.com",
		"jesbass": "jesbass@gmail.com",
		"doug":    "doug@gmail.com",
	}

	// for k, v := range emails {
	// 	fmt.Printf("%s: %s\n", k, v)
	// }

	for k := range emails {
		fmt.Printf("Name: %s\n", k)
	}

	for _, v := range emails {
		fmt.Printf("Email: %s\n", v)
	}

}
