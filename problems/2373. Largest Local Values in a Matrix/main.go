package main

import "fmt"

func largestLocal(grid [][]int) [][]int {
	maxLocal := make([][]int, 0)
	n := len(grid)

	for i := 0; i < n-2; i++ {
		arr := make([]int, 0)
		for j := 0; j < n-2; j++ {
			maxVal := 0
			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					maxVal = max(maxVal, grid[k][l])
				}
			}
			arr = append(arr, maxVal)
		}
		maxLocal = append(maxLocal, arr)
	}
	return maxLocal
}

func main() {
	fmt.Println(largestLocal([][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}))
}
