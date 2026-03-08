package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	// var result strings.Builder
	// for _, c := range address {
	// 	if c == '.' {
	// 		result.WriteString("[.]")
	//		// result += "[.]"
	// 	} else {
	// 		result.WriteRune(c)
	//		result += string(c)
	// 	}
	// }
	// return result.String()
	return strings.ReplaceAll(address, ".", "[.]")
}

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}
