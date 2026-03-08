package main

import (
	"fmt"
	// "slices"
)

// func quickSort(s []int) []int {
// 	var pivot int
// 	less := make([]int, 0)
// 	greater := make([]int, 0)
// 	if len(s) < 2 {
// 		return s
// 	} else {
// 		pivot = s[0]
// 		for _, digit := range s[1:] {
// 			if digit <= pivot {
// 				less = append(less, digit)
// 			} else {
// 				greater = append(greater, digit)
// 			}
// 		}
// 		return append(append(quickSort(less), pivot), quickSort(greater)...)
// 	}
// }

// func sortedSquares(nums []int) []int {
// 	squared := make([]int, 0)
// 	for _, number := range nums {
// 		squared = append(squared, number*number)
// 	}
// return quickSort(squared)
// }

func sortedSquares(nums []int) []int {

	return []int{0}
}

func main() {
	// fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
