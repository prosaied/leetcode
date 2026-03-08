package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	var richest int = 0
	var tempBalanc int = 0
	for _, account := range accounts {
		for _, balance := range account {
			tempBalanc += balance
		}
		if tempBalanc >= richest {
			richest = tempBalanc
		}
		tempBalanc = 0
	}
	return richest
}

func main() {
	// fmt.Println(maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
	// fmt.Println(maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))
	fmt.Println(maximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}))
}
