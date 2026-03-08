package main

import (
	"fmt"
)

func xorOperation(n int, start int) int {
	var numbers []int
	var result int
	for i := range n {
		numbers = append(numbers, (start + (2 * i)))
	}
	for _, number := range numbers {
		result = result ^ number
	}
	return result
}

func main() {
	fmt.Println(xorOperation(5, 0))
	fmt.Println(xorOperation(4, 3))
}
