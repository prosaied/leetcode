package main

import "fmt"

func reverseDegree(s string) int {
	var result int
	for index, c := range s {
		result += (26 - (int(c) - 'a')) * (index + 1)
	}
	return result
}

func main() {
	fmt.Println(reverseDegree("zaza"))
}
