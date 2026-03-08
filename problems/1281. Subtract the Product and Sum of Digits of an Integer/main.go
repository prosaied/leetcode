package main

import "fmt"

func subtractProductAndSum(n int) int {
	var sums, products int = 0, 1
	for n > 0 {
		sums += n % 10
		products = products * (n % 10)
		n = n / 10
	}
	return products - sums
}

func main() {
	fmt.Println(subtractProductAndSum(234))
}
