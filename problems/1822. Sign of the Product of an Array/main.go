package main

import "fmt"

func arraySign(nums []int) int {
	var manfi int
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			manfi++
		}
	}
	if manfi%2 == 0 {
		return 1
	} else {
		return -1
	}
}

func main() {
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1}))
}
