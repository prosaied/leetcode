package main

import (
	"fmt"
	"os"
	"sort"
)

func sortArray(nums []int) []int {
	copied := append([]int(nil), nums...)
	sort.Ints(copied)
	return copied
}

func popFirst(nums []int) (int, []int) {
	return nums[0], nums[1:]
}

func numberGame(nums []int) []int {
	if len(nums)%2 != 0 {
		fmt.Println("Array length is odd. Exiting.")
		os.Exit(1)
	}
	var gameResult []int
	var aliceMove int
	var bobMove int
	sortedArray := sortArray(nums)
	for len(sortedArray) > 0 {
		aliceMove, sortedArray = popFirst(sortedArray)
		bobMove, sortedArray = popFirst(sortedArray)
		gameResult = append(gameResult, bobMove, aliceMove)
	}
	return gameResult
}

func main() {
	fmt.Println(numberGame([]int{5, 4, 2, 3, 7, 8}))
}
