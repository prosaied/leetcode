package main

import "fmt"

func buildArray(nums []int) []int {
	// ans := make([]int, 0)
	// for i := 0; i < len(nums); i++ {
	// 	ans = append(ans, nums[nums[i]])
	// }
	// return ans
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}

func main() {
	fmt.Println(buildArray([]int{0, 2, 1, 5, 3, 4}))
}
