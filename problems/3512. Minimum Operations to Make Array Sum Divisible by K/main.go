package main

import "fmt"

func minOperations(nums []int, k int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum % k
}

func main() {
	fmt.Println(minOperations([]int{3, 9, 7}, 5))
}
