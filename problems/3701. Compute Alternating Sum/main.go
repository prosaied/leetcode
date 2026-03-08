package main

import "fmt"

func alternatingSum(nums []int) int {
	var result int = 0

	for index, num := range nums {
		if index%2 == 0 {
			result = result + num
		} else {
			result = result - num
		}
	}
	return result
}

func main() {
	fmt.Println(alternatingSum([]int{1, 3, 5, 7}))
}
