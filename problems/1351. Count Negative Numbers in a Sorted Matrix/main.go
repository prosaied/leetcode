package main

import "fmt"

func countNegatives(grid [][]int) int {
	result := 0
	for _, gr := range grid {
		for _, g := range gr {
			if g < 0 {
				result++
			}
		}
	}
	return result
}

func main() {
	fmt.Println(countNegatives([][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}))
}
