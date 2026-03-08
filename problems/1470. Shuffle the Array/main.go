package main

import "fmt"

func shuffle(nums []int, n int) []int {
	result := make([]int, 2*n)
	var idx int
	for i := 0; i < n; i++ {
		result[idx] = nums[i]
		result[idx+1] = nums[i+n]
		idx += 2
	}

	return result
}

// func shuffle(nums []int, n int) []int {
//     res := []int{}

//     for i := 0; i < n; i++ {
//         res = append(res, nums[i], nums[i + n])
//     }

//     return res
// }

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}
