package main

import "fmt"

func numberOfMatches(n int) int {
	var games int
	for n > 1 {
		games += (n / 2)
		if n%2 == 0 {
			n = (n / 2)
		} else {
			n = (n / 2) + 1
		}
	}
	return games
}

func main() {
	fmt.Println(numberOfMatches(19))
}

// simplest answer would be return n-1 :)


()()()