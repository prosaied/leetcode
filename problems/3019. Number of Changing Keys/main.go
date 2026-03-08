package main

import (
	"fmt"
	"strings"
)

func countKeyChanges(s string) int {
	typedKeys := strings.ToLower(s)
	var changed int = 0
	for i := 1; i < len(s); i++ {
		if typedKeys[i] != typedKeys[i-1] {
			changed += 1
		}
	}
	return changed
}

func main() {
	fmt.Println(countKeyChanges("aAbBcC"))
	fmt.Println(countKeyChanges("AaAaAaaA"))
}
