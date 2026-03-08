package main

import "fmt"

func binarySearch(friends []int, target, low, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2

	if friends[mid] == target {
		return mid
	} else if friends[mid] < target {
		return binarySearch(friends, target, mid+1, high)
	} else {
		return binarySearch(friends, target, low, mid-1)
	}
}

func recoverOrder(order []int, friends []int) []int {
	out := make([]int, 0)
	for _, o := range order {
		result := binarySearch(friends, o, 0, len(friends)-1)
		if result != -1 {
			out = append(out, friends[result])
		}
	}
	return out
}

func main() {
	fmt.Println(recoverOrder([]int{3, 1, 2, 5, 4}, []int{1, 3, 4}))
	fmt.Println(recoverOrder([]int{1, 4, 5, 3, 2}, []int{2, 5}))
}
