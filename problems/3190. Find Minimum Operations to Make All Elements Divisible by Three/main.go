package main

import "fmt"

func minimumOperations(nums []int) int {
	var result int
	for _, num := range nums {
		if num%3 != 0 {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(minimumOperations([]int{1, 2, 3, 4}))
}
