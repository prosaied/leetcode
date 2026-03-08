package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	var result int
	for _, jc := range jewels {
		for _, sc := range stones {
			if jc == sc {
				result++
			}
		}
	}
	return result
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}
