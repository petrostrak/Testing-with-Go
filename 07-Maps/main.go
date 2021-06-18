package main

import "fmt"

var (
	dictionary = map[string]string{}
)

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

func main() {
	dictionary["test"] = "this is just a test"
	fmt.Println(Search(dictionary, "test"))
}
