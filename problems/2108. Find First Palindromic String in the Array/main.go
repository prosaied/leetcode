package main

import "fmt"

func firstPalindrome(words []string) string {
	// var b []byte
	// var bRev []byte
	// for _, word := range words {
	// 	b = []byte(word)
	// 	for i := len(b) - 1; i >= 0; i-- {
	// 		bRev = append(bRev, b[i])
	// 	}
	// 	if string(bRev) == string(b) {
	// 		return string(b)
	// 	}
	// 	bRev = nil
	// }
	// return ""
	var b []byte
	var reversed string
	for _, word := range words {
		b = []byte(word)
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
		reversed = string(b)
		if reversed == word {
			return word
		}
	}
	return ""
}

func main() {
	words := []string{"abc", "car", "ada", "racecar", "cool"}
	fmt.Println(firstPalindrome(words))
}
