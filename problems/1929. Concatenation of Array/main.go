package main

import "fmt"

func getConcatenation(nums []int) []int {
	ans := make([]int, len(nums)*2)
	for i := 0; i < len(nums); i++ {
		ans[i+len(nums)] = nums[i]
		ans[i] = nums[i]
	}
	return ans
	// return append(nums, nums...)
}

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 1}))
}
