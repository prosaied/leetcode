package main

import "fmt"

func getSneakyNumbers(nums []int) []int {
	// numbersCount := make(map[int]int)
	// result := make([]int, 0)
	// for _, n := range nums {
	// 	numbersCount[n]++
	// }
	// for num, count := range numbersCount {
	// 	if count > 1 {
	// 		result = append(result, num)
	// 	}
	// }
	// return result

	numbersCount := make(map[int]int)
	result := make([]int, 0)

	for _, n := range nums {
		numbersCount[n]++
		if numbersCount[n] > 1 {
			result = append(result, n)
		}
	}
	return result
}

func main() {
	array := []int{5, 4, 7, 11, 23, 34, 56, 34}
	fmt.Println(getSneakyNumbers(array))
}
