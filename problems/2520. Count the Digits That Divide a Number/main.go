package main

import "fmt"

func countDigits(num int) int {
	n := num
	var i int = 0
	for n > 0 {
		digit := n % 10
		if num%digit == 0 {
			i++
		}
		n = n / 10
	}
	return i

}

func main() {
	fmt.Println(countDigits(7))
	fmt.Println(countDigits(121))
	fmt.Println(countDigits(1248))
}
