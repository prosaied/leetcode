package main

import "fmt"

func abs(numb int) int {
	if numb < 0 {
		return -numb
	}
	return numb
}
func scoreOfString(s string) int {
	// var str []int
	// var score int
	// for _, c := range s {
	// 	str = append(str, int(c))
	// }
	// for i := 0; i < len(str)-1; i++ {
	// 	if str[i] > str[i+1] {
	// 		score = score + (str[i] - str[i+1])
	// 	} else {
	// 		score = score + (str[i+1] - str[i])
	// 	}
	// }
	// return score
	var score int
	for i, c := range s {
		if i != 0 {
			score += abs(int(s[i-1]) - int(c))
		}
	}
	return score
}

func main() {
	fmt.Println(scoreOfString("hello"))
}
