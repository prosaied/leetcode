package main

import "fmt"
import "strings"

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func main() {
	fmt.Println(toLowerCase("Hello"))
}
