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

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

func main() {
	d := Dictionary{}
	d.Add("test", "This is just a test")
	fmt.Println(d.Search("test"))
}
