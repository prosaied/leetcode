package main

import "fmt"

func balancedStringSplit(s string) int {
	balance, answer := 0, 0
	for _, c := range s {
		if c == 'R' {
			balance++
		} else {
			balance--
		}
		if balance == 0 {
			answer++
		}
	}
	return answer
}

func main() {
	result := balancedStringSplit("RLRRLLRLRL")
	fmt.Println(result)

}
