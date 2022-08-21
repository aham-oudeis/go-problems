package main

import (
	"fmt"
	"strings"
)

func swapCase(str string) string {
	b := strings.Builder{}

	for idx := range str {
		char := string(str[idx])

		if strings.Contains("abcdefghijklmnopqrstuvwxyz", char) {
			char = strings.ToUpper(char)
		} else {
			char = strings.ToLower(char)
		}

		b.WriteString(char)
	}

	return b.String()
}

func main() {
	fmt.Println(swapCase("CamelCase"))
	fmt.Println("vim-go")
}
