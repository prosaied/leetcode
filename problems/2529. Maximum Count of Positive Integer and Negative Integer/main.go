package main

import "fmt"

func maximumCount(nums []int) int {
	var neg, pos int
	for _, digit := range nums {
		if digit < 0 {
			neg++
		}
		if digit > 0 {
			pos++
		}
	}
	if pos > neg {
		return pos
	}
	return neg
}

func main() {
	fmt.Println(maximumCount([]int{-2, -1, -1, 1, 2, 3}))
}
