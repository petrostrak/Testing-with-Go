package main

import "fmt"

func Repeat(s string) string {
	var outcome string
	for i := 0; i < 5; i++ {
		outcome += s
	}
	return outcome
}

func main() {
	fmt.Println(Repeat("a"))
}
