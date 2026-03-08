package main

import (
	"strings"
)

func checkIfPangram(sentence string) bool {
	loweredSentence := strings.ToLower(sentence)
	alphabet := make(map[rune]bool)
	for _, value := range loweredSentence {
		alphabet[value] = true
	}
	if len(alphabet) == 26 {
		return true
	} else {
		return false
	}
}

func main() {
	checkIfPangram("salam")
}
