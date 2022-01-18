package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign key values
	// emails["roms"] = "roms@gmail.com"
	// emails["jesbass"] = "jesbass@gmail.com"
	// emails["line"] = "line@gmail.com"
	// emails["layne"] = "layne@gmail.com"
	// emails["emy"] = "emy@gmail.com"
	// emails["lelets"] = "lelets@gmail.com"
	// emails["doug"] = "doug@gmail.com"

	// declare map and assign kv
	emails := map[string]string{
		"roms":    "roms@gmail.com",
		"jesbass": "jesbass@gmail.com",
		"doug":    "doug@gmail.com",
	}

	fmt.Println(emails)
	fmt.Println(emails["jesbass"])
	fmt.Println(len(emails))

	// delete from map
	delete(emails, "doug")
	fmt.Println(emails)
}
