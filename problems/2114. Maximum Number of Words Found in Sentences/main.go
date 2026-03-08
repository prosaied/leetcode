package main

import (
	"fmt"
	"unicode"
)

func mostWordsFound(sentences []string) int {
	top := 0
	i := 0
	for _, sentence := range sentences {
		for _, char := range sentence {
			if unicode.IsSpace(char) == true {
				i++
			}
		}
		if i > top {
			top = i
		}
		i = 0
	}
	top++
	return top
}

func main() {
	sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	fmt.Println(mostWordsFound(sentences))
}
