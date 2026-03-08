package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i] > nums[j] {
				result[i]++
			}
		}
	}
	return result
}

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
}

// with mapping it could be less complicated
// map frequency and then compare
