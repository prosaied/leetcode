package main

import "fmt"

func judgeCircle(moves string) bool {
	var location [2]int
	for _, v := range moves {
		switch v {
		case 'R':
			location[0]++
		case 'L':
			location[0]--
		case 'U':
			location[1]++
		case 'D':
			location[1]--
		}
	}
	return location[0] == 0 && location[1] == 0
}

func main() {
	fmt.Println(judgeCircle("UD"))
}
