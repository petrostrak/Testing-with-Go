package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GreetToWriter(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func GreetToBuffer(w *bytes.Buffer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	GreetToWriter(w, "Petros")
}

func main() {
	b := bytes.Buffer{}
	GreetToBuffer(&b, "Petros")
	GreetToWriter(os.Stdout, "Petros")
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(GreetHandler)))
}
