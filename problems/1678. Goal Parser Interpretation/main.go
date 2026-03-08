package main

import (
	"fmt"
	"strings"
)

func interpret(command string) string {
	command1 := strings.ReplaceAll(command, "()", "o")
	command2 := strings.ReplaceAll(command1, "(al)", "al")
	return command2
}

func main() {
	fmt.Println(interpret("G()()()()(al)"))
}
