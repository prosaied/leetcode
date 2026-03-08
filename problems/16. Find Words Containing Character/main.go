package main

import "fmt"

func findWordsContaining(words []string, x byte) []int {
	var result []int
	for index, word := range words {
		for _, w := range word {
			if w == rune(x) {
				result = append(result, index)
				break
			}
		}
	}
	return result
}

func main() {
	fmt.Println(findWordsContaining([]string{"leet", "code"}, 'e'))
	fmt.Println(findWordsContaining([]string{"abc", "bcd", "aaaa", "cbc"}, 'a'))
	fmt.Println(findWordsContaining([]string{"abc", "bcd", "aaaa", "cbc"}, 'z'))
}
