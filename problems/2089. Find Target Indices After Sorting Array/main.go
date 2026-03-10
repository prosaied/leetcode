package main

import (
	"fmt"
)

func sortIndices(nums []int) []int {
	var pivot int
	var less []int
	var greater []int

	if len(nums) < 2 {
		return nums
	} else {
		pivot = nums[0]
		for _, digit := range nums[1:] {
			if digit <= pivot {
				less = append(less, digit)
			} else {
				greater = append(greater, digit)
			}
		}
		return append(append(sortIndices(less), pivot), sortIndices(greater)...)
	}
}

func targetIndices(nums []int, target int) []int {
	var result []int
	for i, digit := range sortIndices(nums) {
		if digit == target {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	fmt.Println(targetIndices([]int{1, 2, 5, 2, 3}, 2))
}
