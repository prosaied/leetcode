package main

import "fmt"

func countConsistentStrings(allowed string, words []string) int {
	var result, counter int
	for _, word := range words {
		counter = 0
		for _, w := range word {
			for _, c := range allowed {
				if w == c {
					counter++
				}
			}
		}
		if counter == len(word) {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
}
