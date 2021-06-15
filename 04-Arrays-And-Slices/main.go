package main

import "fmt"

func Sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}

	return sum
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))
}
