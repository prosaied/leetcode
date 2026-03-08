package main

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	var result []int
	for _, operation := range operations {
		if operation == "C" {
			result = result[:len(result)-1]
		} else if operation == "D" {
			result = append(result, (result[len(result)-1] * 2))
		} else if operation == "+" {
			result = append(result, (result[len(result)-1] + result[len(result)-2]))
		} else {
			i, err := strconv.Atoi(operation)
			if err != nil {
				fmt.Println("not a valid operation")
			}
			result = append(result, i)
		}
	}
	var sum int = 0
	for j := 0; j < len(result); j++ {
		sum = sum + result[j]
	}
	return sum
}

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
}
