package main

import "fmt"

func transformArray(nums []int) []int {
	var even, odd int
	result := make([]int, 0)
	for _, num := range nums {
		if num%2 != 0 {
			odd++
		} else {
			even++
		}
	}
	for i := 0; i < even; i++ {
		result = append(result, 0)
	}
	for i := 0; i < odd; i++ {
		result = append(result, 1)
	}

	return result
}

func main() {
	fmt.Println(transformArray([]int{4, 3, 2, 1}))
}
