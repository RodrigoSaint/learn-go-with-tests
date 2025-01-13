package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writter io.Writer, name string) {
	fmt.Fprintf(writter, "Hello, %s", name)
}

func myGreetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Message received")
	Greet(w, "world")
}

func main() {
	fmt.Println("Listening on localhost:5001")
	http.ListenAndServe(":5001", http.HandlerFunc(myGreetHandler))
}
