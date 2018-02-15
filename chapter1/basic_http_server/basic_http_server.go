package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	port := 8000

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World\n")
}
