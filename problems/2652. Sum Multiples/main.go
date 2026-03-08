package main

import "fmt"

func sumOfMultiples(n int) int {
	var result int = 0
	for number := 1; number <= n; number++ {
		if ((number % 3) == 0) || ((number % 5) == 0) || ((number % 7) == 0) {
			result = result + number
		}
	}
	return result
}

func main() {
	fmt.Println(sumOfMultiples(7))
}
