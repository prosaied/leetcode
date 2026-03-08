package main

import "fmt"

// func numIdenticalPairs(nums []int) int {
// 	var result int
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] == nums[j] {
// 				result++
// 			}
// 		}
// 	}
// 	return result
// }

func numIdenticalPairs(A []int) int {
	ans := 0
	cnt := make(map[int]int)

	for _, x := range A {
		ans = ans + cnt[x]
		cnt[x]++
		fmt.Println("ans is: ", ans, "    cnt is: ", cnt)
	}

	return ans
}

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))

}
