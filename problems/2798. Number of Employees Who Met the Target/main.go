package main

import "fmt"

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	var count int
	for _, hour := range hours {
		if hour >= target {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(numberOfEmployeesWhoMetTarget([]int{0, 1, 2, 3, 4}, 2))
}
