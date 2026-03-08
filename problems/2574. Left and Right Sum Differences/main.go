package main

import "fmt"

func leftRightDifference(nums []int) []int {
	leftsum := []int{}
	rightsum := []int{}
	answer := []int{}

	for i := 0; i < len(nums); i++ {
		leftsum = append(leftsum, sum(nums[:i]))
		rightsum = append(rightsum, sum(nums[i+1:]))
		answer = append(answer, abs(leftsum[i]-rightsum[i]))
	}
	return answer
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sum(nums []int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

func main() {
	fmt.Println(leftRightDifference([]int{10, 4, 8, 3}))
}
