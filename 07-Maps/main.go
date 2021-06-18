package main

import (
	"fmt"
)

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrWordNotFound = DictionaryErr("could not find the word you are looking for")
	ErrWordExists   = DictionaryErr("cannot add word because it already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case ErrWordExists:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func main() {
	d := Dictionary{}
	d.Add("test", "This is just a test")
	fmt.Println(d.Search("test"))
}
