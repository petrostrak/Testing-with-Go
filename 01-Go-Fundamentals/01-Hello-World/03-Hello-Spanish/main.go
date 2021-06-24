package main

import "fmt"

func Hello(name, language string) string {
	if name == "" || language == "" {
		return "Hello, World"
	}

	if language == "spanish" {
		return "Hola, " + name
	}

	return "Hello, " + name
}

// It is good to separate your "domain" code from the outside world (side-effects).
// The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain
func main() {
	fmt.Println(Hello("Petros", "spanish"))
}
