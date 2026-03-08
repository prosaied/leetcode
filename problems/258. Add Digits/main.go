package main

import "fmt"

func addDigits(num int) int {
	number := num
	var sum int = 0
	if number < 10 {
		return number
	}
	for number > 0 {
		digit := number % 10
		sum = digit + sum
		number = number / 10
	}
	return addDigits(sum)
}

func main() {
	fmt.Println(addDigits(123))
	fmt.Println(addDigits(38))
	fmt.Println(addDigits(9))
	fmt.Println(addDigits(999))
	fmt.Println(addDigits(9999))
	fmt.Println(addDigits(1))
}
