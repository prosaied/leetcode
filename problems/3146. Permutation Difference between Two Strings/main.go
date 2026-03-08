package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func findPermutationDifference(s string, t string) int {
	var result int
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)

	for index, character := range s {
		sMap[character] = index
	}
	for index, character := range t {
		tMap[character] = index
	}
	for c := range tMap {
		result += (abs(tMap[c] - sMap[c]))
	}
	return result
}

func main() {
	fmt.Println(findPermutationDifference("abc", "bac"))
}
