package main

import (
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>Hello World </h1>")
}

func about(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>About Page </h1>")
}

func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server started running at port :3000")
	http.ListenAndServe(":3000", nil)
}
