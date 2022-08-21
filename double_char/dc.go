package main

import (
	"fmt"
	"regexp"
)

func repeater(str string) string {
	pattern := regexp.MustCompile(`(.)`)

	return pattern.ReplaceAllString(str, "$1$1")
}

func doubleConsonants(str string) string {
	pattern := regexp.MustCompile(`(b|c|d|f|g|h|j|k|l|m|n|p|q|r|s|t|v|w|x|y|z])`)

	return pattern.ReplaceAllString(str, "$1$1")
}
func main() {
	fmt.Println(repeater("hello"))
	fmt.Println(doubleConsonants("hello"))
	fmt.Println("vim-go")
}
