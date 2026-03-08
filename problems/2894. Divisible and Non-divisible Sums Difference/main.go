package main

import "fmt"

func differenceOfSums(n int, m int) int {
	var num1 int = 0
	var num2 int = 0
	for i := 1; i <= n; i++ {
		if i%m != 0 {
			num1 += i
		}
		if i%m == 0 {
			num2 += i
		}
	}
	return num1 - num2
}

func main() {
	fmt.Println(differenceOfSums(10, 3))
	fmt.Println(differenceOfSums(5, 1))
}
