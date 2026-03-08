package main

import "fmt"

func canAliceWin(nums []int) bool {
	var singles int
	var doubles int

	for _, num := range nums {
		if num < 10 {
			singles += num
		} else if num >= 10 {
			doubles += num
		}
	}
	if singles != doubles {
		return true
	}
	return false
}

func main() {
	fmt.Println(canAliceWin([]int{1, 2, 3, 4, 10}))
}
