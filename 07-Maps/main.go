package main

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

var (
	ErrWordNotFound = errors.New("could not find the word you are looking for")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil
}

func main() {
	d := Dictionary{"test": "This is just a test"}
	fmt.Println(d.Search("test"))
}
