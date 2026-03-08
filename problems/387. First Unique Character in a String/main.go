package main

import "fmt"

func firstUniqChar(s string) int {
	characters := make(map[rune]int)
	for _, char := range s {
		characters[char]++
	}
	for index, charact := range s {
		if characters[charact] == 1 {
			return index
		}
	}
	return -1

}

func main() {
	fmt.Println(firstUniqChar("abcab"))
	fmt.Println(firstUniqChar("saeed"))
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("aabb"))

}
