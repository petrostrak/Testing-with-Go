package main

import "fmt"

func Sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}

	return sum
}

func SumAll(slices ...[]int) []int {
	var sum []int
	for _, slice := range slices {
		sum = append(sum, Sum(slice))
	}

	return sum
}

func main() {
	fmt.Println(Sum([]int{1, 2}))
	fmt.Println(SumAll([]int{1, 2}, []int{0, 9}))
}
