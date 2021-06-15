package main

import "fmt"

func Hello(name string) string {
	if name == "" {
		return "Hello, World"
	}
	return "Hello, " + name
}

// It is good to separate your "domain" code from the outside world (side-effects).
// The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain
func main() {
	fmt.Println(Hello("Petros"))
}
