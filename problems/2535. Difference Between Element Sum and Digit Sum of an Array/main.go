package main

import (
	"fmt"
)

func differenceOfSum(nums []int) int {
	var elementSum int
	var digitSum int
	var sum int
	for _, num := range nums {
		elementSum = elementSum + num
		for num > 0 {
			sum = sum + (num % 10)
			num = num / 10
		}
		digitSum = digitSum + sum
		num = 0
		sum = 0
	}
	if elementSum > digitSum {
		return elementSum - digitSum
	} else {
		return digitSum - elementSum
	}
}

func main() {
	fmt.Println(differenceOfSum([]int{1, 15, 6, 3}))
}
