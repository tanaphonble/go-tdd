package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet - Greeting
func Greet(writter io.Writer, name string) {
	fmt.Fprintf(writter, "Hello, %s", name)
}

// MyGreeterHandler - Just a handler
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
