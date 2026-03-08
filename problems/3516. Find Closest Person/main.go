package main

import (
	"fmt"
)

func absint(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func findClosest(x int, y int, z int) int {
	if absint(x-z) == absint(y-z) {
		return 0
	} else if absint(x-z) > absint(y-z) {
		return 2
	} else {
		return 1
	}

}

func main() {
	fmt.Println(findClosest(2, 7, 4))
}
