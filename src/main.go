package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, world!")
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println("Server error: ", err)
	}
}
