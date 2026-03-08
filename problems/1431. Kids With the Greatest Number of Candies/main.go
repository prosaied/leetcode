package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var result []bool
	var highest int
	for _, candycount := range candies {
		if candycount > highest {
			highest = candycount
		}
	}
	for _, candy := range candies {
		if candy+extraCandies >= highest {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}
