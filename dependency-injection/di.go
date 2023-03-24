package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// We make use of Dependency Injection (injected io.Writer) to test something which was not possible otherwise
// Dependency Injection is basically useful here because we can inject something we have control on.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello! %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World!")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
	// Greet(os.Stdout, "Shikhar")
}
