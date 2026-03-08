package main

import "fmt"

func countPairs(nums []int, target int) int {
	var result int
	for j := 0; j < len(nums); j++ {
		for i := 0; i < j; i++ {
			if nums[i]+nums[j] < target {
				result++
			}
		}
	}
	return result
}

func main() {
	fmt.Println(countPairs([]int{-6, 2, 5, -2, -7, -1, 3}, -2))
}
