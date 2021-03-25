package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/mytest", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("Hello World")
		writer.Write(([]byte)("this is a test page"))
	})
	http.ListenAndServe(":8080", nil)
}
