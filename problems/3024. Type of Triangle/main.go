package main

import "fmt"

func triangleType(nums []int) string {
	if nums[0]+nums[1] <= nums[2] || nums[0]+nums[2] <= nums[1] || nums[1]+nums[2] <= nums[0] {
		return "none"
	} else if nums[0] == nums[1] && nums[0] == nums[2] {
		return "equilateral"
	} else if nums[0] == nums[1] || nums[0] == nums[2] || nums[1] == nums[2] {
		return "isosceles"
	} else {
		return "scalene"
	}
}

func main() {
	fmt.Println(triangleType([]int{5, 3, 8}))
}
