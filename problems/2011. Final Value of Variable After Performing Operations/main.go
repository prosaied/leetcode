package main

import "fmt"

func finalValueAfterOperations(operations []string) int {
	var x int
	for _, op := range operations {
		switch op {
		case "--X":
			x--
		case "X--":
			x--
		case "++X":
			x++
		case "X++":
			x++
		}
	}
	return x
}

func main() {
	fmt.Println(finalValueAfterOperations([]string{"--X", "X++", "X++"}))
}
